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
	Step	int `json:"step"`			//执行的步骤
	ErrorMsg string `json:"error_msg"`  //执行出错的原因
}

// ServerItem 单个任务需要的结构
type ServerItem struct{
	API string `json:"api" binding:"required"`
	Try string `json:"try" binding:"required"`
	Confirm string `json:"confirm" binding:"required"`
	Cancel string `json:"cancel" binding:"required"`
	Status string `json:"status"`						//单位操作是否执行成功
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

	if req.ExecNum >= MAX_EXEC_NUM {
		fmt.Println("已超过最大执行次数")
		return 
	}

	// 根据不同的执行步骤进行操作
	switch req.Step {
	case 0:
		combineTry(&req)
	case 1:

	}
}



// 同步执行 try 步骤
func combineTry(req *ServerForm){
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
		logs.Info("try 操作成功, ID:", req.ID)
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
	return respBody{
		API: url,
		Status: status,
		Body: string(body),
		Error: errStr,
	}
}

// 同步执行 commit 步骤
func combineCommit(req *ServerForm){

}

// 同步执行 cancel 步骤?
func combineCancel(req *ServerForm){

}


func (rq *ServerForm) toString() string{
	jsonByte, _ := json.Marshal(rq)
	return string(jsonByte)
}


// 并行操作 try 不通过直接取消事务,
func (rq *ServerForm) cancel(errMsg []respBody){
	logs.Info("准备开始取消",rq)
	errStr := JSONToStr(errMsg)
	db, _ := sql.Open("mysql", "root:123123@tcp(127.0.0.1:33060)/go?charset=utf8")
	sql := "UPDATE rpc_ts SET payload=?, status=2,exec_num=exec_num+1 ,update_at=?,error_info=? WHERE id=?"
	stmt, err := db.Prepare(sql)
	checkErr(err)
	_, err = stmt.Exec(rq.toString(),time.Now().Unix(),errStr,rq.ID)
	checkErr(err)
	logs.Debug("插入数据库内容:",rq.toString(),time.Now().Unix(),errStr,rq.ID)


	logs.Debug("此处应该向某处发送错误通知")
}



// ==================================TOOLS===================================

// JSONToStr 转字符串的统一方法
func JSONToStr(req interface{}) string{
	strByte, _ := json.Marshal(req)
	return string(strByte)
}