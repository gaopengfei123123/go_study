package services

import(
	"fmt"
	"encoding/json"
)

const(	
	SYNC_MODE = "sync"		//同步操作模式
	ASYNC_MODE = "async"		//异步操作模式
)

// ServerService 用于处理队列任务的模块
func ServerService(jsonStr []byte){
	requestForm := ClientForm{}
	json.Unmarshal(jsonStr, &requestForm)
	
	if requestForm.Type == SYNC_MODE {
		syncHandler(requestForm)
	}

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
func syncHandler(req ClientForm){
	fmt.Println("这里是同步操作")
}

func asyncHandler(req ClientForm){
	fmt.Println("这里是异步操作")
}