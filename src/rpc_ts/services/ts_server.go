package services

import(
	"fmt"
	"encoding/json"
	// "sync"
	"context"
	"time"
	"net/http"
	"io/ioutil"
	"bytes"
	"github.com/astaxie/beego/logs"
	// 引入 mysql 驱动
	"database/sql"
	_ "github.com/GO-SQL-Driver/MySQL"
)

const(	
	SYNC_MODE = "sync"		//同步操作模式
	ASYNC_MODE = "async"		//异步操作模式
)

func init(){
	fmt.Println("初始化 log 配置")
	// log 开异步
	logs.Async(1e3)
	config := fmt.Sprintf(`{"filename":"%s","separate":["error", "warning", "notice", "info", "debug"]}`, LOG_PATH )
	logs.SetLogger(logs.AdapterConsole, config)
}


// ServerForm 接收参数时的 json 格式
type ServerForm struct{
	Type string `json:"type" binding:"required"`
	Task []ServerItem
	ExecNum int `json:"exec_num"`		//执行次数
	ExecTime int `json:"exec_time"`		//执行时间
	ID	int 	`json:"ID"`				//数据库主键
	Step	int `json:"step"`			//执行的步骤 0为未开始, 1时当为并发时已完成 try,当是串行时再商量
	ErrorMsg string `json:"error_msg"`  //执行出错的原因
}

// ServerItem 单个任务需要的结构
type ServerItem struct{
	API string `json:"api" binding:"required"`
	Try string `json:"try" binding:"required"`
	Confirm string `json:"confirm" binding:"required"`
	Cancel string `json:"cancel" binding:"required"`
	Step int `json:"step"`					//单位操作执行阶段 0 为 try 阶段,1为 commit 阶段,2为 cancel 阶段
	TryStatus string	`json:"try_status"`		// 各阶段的执行任务 有 wait,done,false 三种情况
	CommitStatus string `json:"commit_status"`
	CancelStatus string `json:"cancel_status"`
}

// ServerService 用于处理队列任务的模块
func ServerService(jsonStr []byte){
	requestForm := ServerForm{}
	json.Unmarshal(jsonStr, &requestForm)

	switch requestForm.Type {
	case SYNC_MODE:
		syncHandler(requestForm)
	case ASYNC_MODE:
		asyncHandler(requestForm)
	default:
		fmt.Println("操作异常")
	}
}



// 顺序执行
func syncHandler(req ServerForm){
	fmt.Println("这里是同步操作", req)
}

// 并发执行
func asyncHandler(req ServerForm){
	fmt.Println("这里是异步操作")

	// 检测执行次数
	if req.ExecNum >= MAX_EXEC_NUM {
		fmt.Println("已超过最大执行次数")
		return 
	}
	// 执行次数加1
	req.ExecNum ++

	// 根据不同的执行步骤进行操作
	switch req.Step {
	case 0:
		req.combineTry()
	case 1:

	}
}



// 同步执行 try 步骤
func (req *ServerForm) combineTry(){
	// 创建对应数量的通信通道接收消息
	resChan := make(chan respBody, len(req.Task))
	defer close(resChan)

	logs.Info("执行 try,ID:", req.ID)


	for index, value := range req.Task{
		go execItem("try", index, value, resChan)
	}

	// 判断整体的 try 是否通过
	isPass := true
	// 将异步的数据导出
	var result []respBody 
	for i:= 0; i < len(req.Task); i++ {
		res := <-resChan
		
		if res.Status != 200 {
			isPass = false
		}

		result = append(result, res)
	}

	logs.Info("try 操作完成,ID:", req.ID, " 返回结果:", JSONToStr(result))
	
	// 如果不通过的话
	if !isPass {
		req.cancel(result)
	} else {
		req.Step ++
		logs.Debug("try 操作成功, ID:", req.ID,req.Step)
		// 更新数据库信息
		req.updateStatus()
		// 执行并发提交
		req.combineCommit()
	}
}

// 执行单个请求操作, 可根据 taskType 来区分请求的包体
func execItem(taskType string, index int,task ServerItem,  resultChan chan<- respBody){
	// 类似 try catch 的一个放报错机制
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
		}
	}()
	// defer wg.Done()

	// 创建一个阻塞通道,用于执行请求任务,也是为了计算超时时间
	dst := make(chan respBody)
	defer close(dst)

	// 这里是进行 post 请求的主体
	go func(task ServerItem, taskType string) {
		start := time.Now()
		var resp respBody
		var url string
		switch taskType {
		case "try":
			url = fmt.Sprintf("%s/try", task.API)
			resp = postClien(url, task.Try)
		case "confirm":
			url = fmt.Sprintf("%s/confirm", task.API)
			resp = postClien(url, task.Confirm)
		case "cancel":
			url = fmt.Sprintf("%s/cancel", task.API)
			resp = postClien(url, task.Cancel)
		}
		// 对返回内容打上 api 信息
		resp.API = url
		
		end := time.Now()
		exeTime := end.Sub(start).Nanoseconds() / 1000000

		logs.Info(task.API, ":"," delay:" ,  exeTime , "ms",resp.Status, string(resp.Body))

		// 将请求结果导出
		dst <- resp

	}(task, taskType)

	// 这里就是声明一个倒计时
	ctx, cancel := context.WithTimeout(context.Background(),MAX_EXEC_TIME * time.Second)
	defer cancel()

	// 监听超时时间
	LOOP:
	for {
		select {
		case resp := <-dst:
			// 当请求正常时直接怼到结果信道中
			resp.Index = index
			resultChan <- resp
			break LOOP
		case <-ctx.Done():
			// 当请求超时时,需要生成 log 并返回错误信息
			errResp := respBody{
				API: task.API,
				Status: 400,
				Body: "",
				Error: "exec timeout",
				Index: index,
				ErrorCode: 408,
			}
			resultChan <- errResp

			logs.Warn("URL: %s has been running toooooooo looooooong! \n",task.API)
			break LOOP
		}
	}

}
// 同步执行 commit 步骤
// 思路:
// 1. 首先过滤出来未执行和执行失败的任务下标,准备进行处理
// 2. 进行批处理执行
// 3. 根据执行请求的返回来进行操作: 1>如果成功那么做标记后返回内容,2> 如果失败: 一如果出现执行超时情况则标记失败准备下次请求,如果是接口返回无法提交或500则整体执行 cancel 操作
func (req *ServerForm) combineCommit(){
	// 首先筛选出来未执行的任务的下标, 因为执行到这一步的时候已经是 try 操作完毕了
	// 因此这里主要看的是 ServerForm 的Step 和 Item 当中的 commit 的执行状态
	logs.Debug("执行 commit 逻辑 ","执行步骤:", req.Step)

	resChan := make(chan respBody, len(req.Task))
	defer close(resChan)

	// 装载合法数据的索引, 标明哪几条任务再执行,然后也是接收这几条任务的 channel,以免发生漏消息或者阻塞
	var indexFilter []int
	for index, value := range req.Task {
		if value.CommitStatus != "done" {
			indexFilter = append(indexFilter, index)
			// 单个执行提交操作
			go execItem("confirm", index, value, resChan)
		}
	}

	// 判断整体的 confirm 是否通过(如果部分未通过则会重回队列)
	isPass := true
	// 判断是否产生需要终止整个事务的事件
	isBreak := false
	// 将异步的数据导出
	var result []respBody 

	for i:= 0; i < len(indexFilter); i++ {
		res := <-resChan
		if res.Status != 200 {
			// 只要请求不成功就不能通过,但是还需要判断是否终止操作
			isPass = false
			// 执行超时的则重回队列
			if res.ErrorCode != 408 {
				isBreak = true
			}
			req.Task[res.Index].CommitStatus = "false"
		} else {
			req.Task[res.Index].CommitStatus = "done"
		}
		result = append(result, res)
	}

	logs.Error("当前事务状态:", req)

	if isBreak {
		logs.Info("执行中断操作操作,ID:", req.ID)
		//  进入下一阶段
		req.Step++
		req.updateStatus()
		req.combineBreak()
		return
	}

	if !isPass {
		logs.Debug("执行重回队列的操作")
		return
	}

	// 完成动作
	logs.Info("准备完成动作")
	req.success()

	

	logs.Debug("当前准备执行的任务索引:", indexFilter)


}

// 执行单个 commit 操作
func execCommit(task ServerItem, resChan chan<- respBody, ){
	// 类似 try catch 的一个放报错机制
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
		}
	}()

	// 创建一个阻塞通道,用于执行请求任务,也是为了计算超时时间
	dst := make(chan respBody)
	defer close(dst)
}

// 同步执行 cancel 步骤?
func combineCancel(req *ServerForm){

}


func (req *ServerForm) toString() string{
	jsonByte, _ := json.Marshal(req)
	return string(jsonByte)
}


// 更新当前任务的mysql状态
func (req *ServerForm) updateStatus(){
	logs.Debug("持久化状态,ID:", req.ID)
	db, _ := sql.Open("mysql", "root:123123@tcp(127.0.0.1:33060)/go?charset=utf8")
	sql := "UPDATE rpc_ts SET payload=?,status=1,exec_num=?,update_at=?,error_info=? WHERE id=?"
	stmt, err := db.Prepare(sql)
	checkErr(err)

	_, err = stmt.Exec(req.toString(), req.ExecNum, time.Now().Unix(), req.ErrorMsg, req.ID)
	checkErr(err)
	logs.Debug("更新数据库信息,ID:", req.ID)
}

// 并行操作 try 不通过直接取消事务,
func (req *ServerForm) cancel(errMsg []respBody){
	logs.Error("准备开始取消")
	errStr := JSONToStr(errMsg)
	db, _ := sql.Open("mysql", "root:123123@tcp(127.0.0.1:33060)/go?charset=utf8")
	sql := "UPDATE rpc_ts SET payload=?, status=20, exec_num=?, update_at=?, error_info=? WHERE id=?"
	stmt, err := db.Prepare(sql)
	checkErr(err)
	_, err = stmt.Exec(req.toString(), req.ExecNum, time.Now().Unix(), errStr, req.ID)
	checkErr(err)
	logs.Debug("取消事务时插入数据库内容:",req.toString(),time.Now().Unix(),errStr,req.ID)
	logs.Debug("此处应该向某处发送 [事务取消] 通知")
}

// 存在非正常取消操作
func (req *ServerForm) crash(errMsg []respBody){
	errStr := JSONToStr(errMsg)
	db, _ := sql.Open("mysql", "root:123123@tcp(127.0.0.1:33060)/go?charset=utf8")
	sql := "UPDATE rpc_ts SET payload=?, status=21, exec_num=?, update_at=?, error_info=? WHERE id=?"
	stmt, err := db.Prepare(sql)
	checkErr(err)
	_, err = stmt.Exec(req.toString(), req.ExecNum, time.Now().Unix(), errStr, req.ID)
	checkErr(err)
	logs.Debug("此处应该向某处发送 [事务异常] 通知")
}

// 事务执行完成动作
func (req *ServerForm) success(){
	db, _ := sql.Open("mysql", "root:123123@tcp(127.0.0.1:33060)/go?charset=utf8")
	sql := "UPDATE rpc_ts SET payload=?, status=2, exec_num=?, update_at=? WHERE id=?"
	stmt, err := db.Prepare(sql)
	checkErr(err)
	_, err = stmt.Exec(req.toString(), req.ExecNum, time.Now().Unix(), req.ID)
	checkErr(err)
	logs.Debug("执行成功的操作完成")
}

// 插入MQ
func (req *ServerForm) insertMQ() {
	jsonStr := req.toString()

	insertKey := fmt.Sprintf("ts_queue_%v_%v", req.ID, req.ExecNum)
	// 向消息队列中发送消息
	var mq MQService
	mq.Send(insertKey,jsonStr)
}

// 执行中断事务的操作
 func (req *ServerForm) combineBreak(){
	logs.Debug("打印当前中断状态时的事务状态", req.toString())

	resChan := make(chan respBody, len(req.Task))
	defer close(resChan)

	// 装载合法数据的索引, 标明哪几条任务再执行,然后也是接收这几条任务的 channel,以免发生漏消息或者阻塞
	var indexFilter []int
	for index, value := range req.Task {
		if value.CommitStatus == "done" {
			indexFilter = append(indexFilter, index)
			// 单个执行取消操作
			go execItem("cancel", index, value, resChan)
		}
	}

	// 将异步的数据导出
	var result []respBody 
	// 判断整体的 cancel 是否通过(如果部分未通过则会重回队列)
	isPass := true
	// 判断是否产生需要终止整个事务的事件
	isBreak := false
	for i:= 0; i < len(indexFilter); i++ {
		res := <-resChan
		if res.Status != 200 {
			// 当执行cancel 都出现错误的时候说明回滚不成功,
			isPass = false
			// 执行超时的则重回队列
			if res.ErrorCode != 408 {
				isBreak = true
			}
			req.Task[res.Index].CancelStatus = "false"
		} else {
			req.Task[res.Index].CancelStatus = "done"
		}
		result = append(result,res)
	}

	if isBreak {
		// 当存在事务无法回滚的情况
		req.crash(result)
		return
	}

	if !isPass {
		// 执行重回队列操作
		logs.Debug("执行重回队列的操作 cancel step")
		req.insertMQ()
		return
	}

	// 全部正常回滚则关闭事务
	req.cancel(result)
	logs.Info("终止操作完毕!ID:", req.ID)
 }


// ==================================TOOLS===================================

// JSONToStr 转字符串的统一方法
func JSONToStr(req interface{}) string{
	strByte, _ := json.Marshal(req)
	return string(strByte)
}


// 统一的返回格式内容
// ```json
// {
// 	"code": 200,	
// 	"error_code": 400,
// 	"error_message": "错误信息",
// 	"data": "返回内容的 json 字符串"
// }

// ```

// 进行 post 请求时的返回的结构体
type respBody struct{
	Status int		`json:"code"`
	Body string		`json:"data"`
	Error string	`json:"error_message"`
	API string
	Index int	// 任务所属的下标
	ErrorCode int `json:"error_code"`
}
// post 请求工具
func postClien(url string, jsonStr string) respBody {
	// 类似 try catch 的一个放报错机制
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
		}
	}()

	var jsonByte = []byte(jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()


	body, _ := ioutil.ReadAll(resp.Body)

	logs.Debug(url, "post 请求返回内容:", string(body))

	if resp.StatusCode == 200 {
		respForm := respBody{}
		json.Unmarshal(body, &respForm)

		if (respForm.ErrorCode != 0) {
			respForm.Status = 401
		}
		
		return respForm
	} 
	// 非200的请求统统属于
	// 返回一个指定的响应结构
	return respBody{
		API: url,
		Status: resp.StatusCode,
		Body: string(body),
		Error: "request error",
		ErrorCode: resp.StatusCode,
	}
	
}