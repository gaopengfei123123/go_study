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

	logs.Info("任务ID:", req.ID, "执行 try")


	for _, value := range req.Task{
		fmt.Println("开始执行任务:",value.API)
		go execTry(value, resChan)
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

	fmt.Println("同步 try", result, "是否通过:", isPass)
}

// 执行单个 try 操作
func execTry(task ServerItem,  resultChan chan<- respBody){
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
	go func(task ServerItem) {
		start := time.Now()
		resp := postClien(task.API, task.Try)
		end := time.Now()
		exeTime := end.Sub(start).Nanoseconds() / 1000000

		logs.Info(task.API, ":"," delay:" ,  exeTime , "ms",resp.Status, string(resp.Body))

		// 将请求结果导出
		dst <- resp

	}(task)

	// 这里就是声明一个倒计时
	ctx, cancel := context.WithTimeout(context.Background(),MAX_EXEC_TIME * time.Second)
	defer cancel()

	// 监听超时时间
	LOOP:
	for {
		select {
		case resp := <-dst:
			// 当请求正常时直接怼到结果信道中
			resultChan <- resp
			break LOOP
		case <-ctx.Done():
			// 当请求超时时,需要生成 log 并返回错误信息
			errResp := respBody{
				API: task.API,
				Status: 400,
				Body: "",
				Error: "exec timeout",
			}
			resultChan <- errResp

			logs.Warn("URL: %s has been running too looong! \n",task.API)
			break LOOP
		}
	}

}


// 进行 post 请求时的返回的结构体
type respBody struct{
	Status int
	Body string
	Error string
	API string
	Index int	// 任务所属的下标
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
	var status int
	var errStr string
	if resp.Status == "200 OK" {
		status = 200
	} else {
		status = 400
		errStr = string(body)
	}
	// 返回一个指定的响应结构
	return respBody{
		API: url,
		Status: status,
		Body: string(body),
		Error: errStr,
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

	// resChan := make(chan respBody, len(req.Task))
	// defer close(resChan)

	var indexFilter []int
	for index, value := range req.Task {
		if value.CommitStatus != "done" {
			indexFilter = append(indexFilter, index)
			// 单个执行提交操作
			// go execCommit(value, resChan, index)
		}
	}

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
	logs.Info("准备开始取消",req)
	errStr := JSONToStr(errMsg)
	db, _ := sql.Open("mysql", "root:123123@tcp(127.0.0.1:33060)/go?charset=utf8")
	sql := "UPDATE rpc_ts SET payload=?, status=2,exec_num=? ,update_at=?,error_info=? WHERE id=?"
	stmt, err := db.Prepare(sql)
	checkErr(err)
	_, err = stmt.Exec(req.toString(), req.ExecNum, time.Now().Unix(), errStr, req.ID)
	checkErr(err)
	logs.Debug("插入数据库内容:",req.toString(),time.Now().Unix(),errStr,req.ID)


	logs.Debug("此处应该向某处发送错误通知")
}



// ==================================TOOLS===================================

// JSONToStr 转字符串的统一方法
func JSONToStr(req interface{}) string{
	strByte, _ := json.Marshal(req)
	return string(strByte)
}