package services

import(
	"fmt"
	"encoding/json"
	"sync"
	"context"
	"time"
	"net/http"
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



// 并行
func syncHandler(req ServerForm){
	fmt.Println("这里是同步操作", req)
}

func asyncHandler(req ServerForm){
	fmt.Println("这里是异步操作")

	if req.ExecNum >= MAX_EXEC_NUM {
		fmt.Println("已超过最大执行次数")
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
	wg := sync.WaitGroup{}

	for _, value := range req.Task{
		fmt.Println("执行任务:",value)
		wg.Add(1)
		go execTry(value, &wg)
	}

	wg.Wait()

	fmt.Println("同步 try")
}

// 执行单个 try 操作
func execTry(task ServerItem, wg *sync.WaitGroup){
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer wg.Done()

	dst := make(chan struct{})
	defer close(dst)

	go func(task ServerItem) {
		start := time.Now()

		resp, err := http.Head(task.API)
		if err != nil {
			fmt.Println("Error:",task.API, err)
		}
		end := time.Now()
		exeTime := end.Sub(start).Nanoseconds() / 1000000
		fmt.Println(task.API, ":", resp.Status, " delay:" ,  exeTime , "ms")

		dst <- struct{}{}

	}(task)

	// 这里就是声明一个倒计时
	ctx, cancel := context.WithTimeout(context.Background(),MAX_EXEC_TIME * time.Second)
	defer cancel()

	LOOP:
	for {
		select {
		case <-dst:
			break LOOP
		case <-ctx.Done():
			fmt.Printf("URL: %s has been running too looong! \n",task.API)
			break LOOP
		}
	}



}

// 同步执行 commit 步骤
func combineCommit(req *ServerForm){

}

// 同步执行 cancel 步骤?
func combineCancel(req *ServerForm){

}