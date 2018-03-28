package services

import(
	"fmt"
	"encoding/json"
	"time"

	// 引入 mysql 驱动
	"database/sql"
	_ "github.com/GO-SQL-Driver/MySQL"
	"github.com/astaxie/beego/logs"
)

func init(){
	fmt.Println("初始化 log 配置")
	// log 开异步
	logs.Async(1e3)
	config := fmt.Sprintf(`{"filename":"%s","separate":["error", "warning", "notice", "info", "debug"]}`, LOG_PATH )
	logs.SetLogger(logs.AdapterMultiFile, config)
}

// ClientForm 接收参数时的 json 格式
type ClientForm struct{
	Type string `json:"type" binding:"required"`
	Task []TaskItem
	ID	 int64
}

// TaskItem 单个任务需要的结构
type TaskItem struct{
	API string `json:"api" binding:"required"`
	Try string `json:"try" binding:"required"`
	Confirm string `json:"confirm" binding:"required"`
	Cancel string `json:"cancel" binding:"required"`
}



// struct 转成 json 字符串
func (cf *ClientForm) toString() string {
	jsonByte, _ := json.Marshal(cf)
	return string(jsonByte)
}

// 插入数据库
func (cf *ClientForm) insertSQL() int64 {
	db, err := sql.Open("mysql", "root:123123@tcp(127.0.0.1:33060)/go?charset=utf8")
	defer db.Close()
	checkErr(err)

	//insert
	stmt, err := db.Prepare("INSERT rpc_ts SET payload=? , status=?,exec_num=?, update_at=? , create_at=?")
	checkErr(err)

	jsonStr := cf.toString()
	timeStamp := time.Now().Unix()
	res, err := stmt.Exec(jsonStr, 0,0,timeStamp,timeStamp )
	checkErr(err)
	// //获取插入数据的 id
	id, err := res.LastInsertId()
	checkErr(err)

	cf.ID = id
	return id
}

// 插入MQ
func (cf *ClientForm) insertMQ() {
	jsonStr := cf.toString()

	insertKey := fmt.Sprintf("ts_queue_%v", cf.ID)
	// 向消息队列中发送消息
	var mq MQService
	mq.Send(insertKey,jsonStr)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}



// Response 通用的返回接口
type Response map[string]interface{}

// ClientService 客户端的运行逻辑
func ClientService(request ClientForm) Response{
	request.insertSQL()
	request.insertMQ()

	return Response{
		"message" : request.Type,
		"api": request.Task[0].API,
		"str": request.toString(),
	}
}


