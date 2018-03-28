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
)

const(	
	SYNC_MODE = "sync"		//同步操作模式
	ASYNC_MODE = "async"		//异步操作模式
)


// ServerForm 接收参数时的 json 格式
type ServerForm struct{
	Type string `json:"type" binding:"required"`
	Task []ServerItem
	ExecNum int `json:"exec_num"`		//执行次数
	ExecTime int `json:"exec_time"`		//执行时间
	ID	int 	`json:"ID"`				//数据库主键
	Step	int `json:"step"`			//执行的步骤
}

// ServerItem 单个任务需要的结构
type ServerItem struct{
	API string `json:"api" binding:"required"`
	Try string `json:"try" binding:"required"`
	Confirm string `json:"confirm" binding:"required"`
	Cancel string `json:"cancel" binding:"required"`
	Status string `json:"status"`
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

	for _, value := range req.Task{
		fmt.Println("开始执行任务:",value)
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
		fmt.Println("执行失败需要处理的步骤")
	}

	fmt.Println("同步 try", result, "是否通过:", isPass)
}

// 执行单个 try 操作
func execTry(task ServerItem,  resultChan chan<- respBody){
	// 类似 try catch 的一个放报错机制
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
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

		fmt.Println(task.API, ":"," delay:" ,  exeTime , "ms",resp.Status, string(resp.Body))

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
				Status: 400,
				Body: "",
				Error: "exec timeout",
			}
			resultChan <- errResp

			fmt.Printf("URL: %s has been running too looong! \n",task.API)
			break LOOP
		}
	}

}


// 进行 post 请求时的返回的结构体
type respBody struct{
	Status int
	Body string
	Error string
}
// post 请求工具
func postClien(url string, jsonStr string) respBody {
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
	if resp.Status == "200 OK" {
		status = 200
	} else {
		status = 400
	}
	return respBody{
		Status: status,
		Body: string(body),
	}
}

// 同步执行 commit 步骤
func combineCommit(req *ServerForm){

}

// 同步执行 cancel 步骤?
func combineCancel(req *ServerForm){

}