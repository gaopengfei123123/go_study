package main

import(
	"github.com/astaxie/beego/logs"
)

func main(){
	logs.SetLogger(logs.AdapterFile, `{"filename":"test.log"}`)
	//an official log.Logger
    l := logs.GetLogger()
    l.Println("this is a message of http")
    //an official log.Logger with prefix ORM
    logs.GetLogger("ORM").Println("this is a message of orm")

    logs.Debug("my book is bought in the year of ", 2016)
    logs.Info("this %s cat is %v years old", "yellow", 3)
    logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
    logs.Error(1024, "is a very", "good game")
    logs.Critical("oh,crash")
}