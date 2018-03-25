package services

import(
	// "fmt"
	"encoding/json"
)

// ClientForm 接收参数时的 json 格式
type ClientForm struct{
	Type string `json:"type" binding:"required"`
	Task []TaskItem
}
func (cf *ClientForm) toString() string {
	jsonByte, _ := json.Marshal(cf)
	return string(jsonByte)
}


// TaskItem 单个任务需要的结构
type TaskItem struct{
	API string `json:"api" binding:"required"`
	Try string `json:"try" binding:"required"`
	Confirm string `json:"confirm" binding:"required"`
	Cancel string `json:"cancel" binding:"required"`
}

// Response 通用的返回接口
type Response map[string]interface{}


// ClientService 客户端的运行逻辑
func ClientService(request ClientForm) Response{
	jsonStr := request.toString()

	// var testJson ClientForm
	// json.Unmarshal([]byte(jsonStr), &testJson)
	// fmt.Println(testJson)

	return Response{
		"message" : request.Type,
		"api": request.Task[0].API,
		"str": jsonStr,
	}
}


