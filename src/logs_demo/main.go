package main

import(
	logs "logs_demo/loghandler"
	"fmt"
	"time"
	"math/rand"
)

type testInfo struct{
	Msg string `json:"msg"`
	Code int `json:"code"`
}

func main(){
	t1 := time.Now()
	info := testInfo{
		"this is msg",
		200,
	}
	for i:=0;i<10;i++ {
		logs.SetUniqueID(rand.Int63n(100))
		logs.Debug(info)
		logs.Info(info)
		logs.Warn(info)
		logs.Critical(info)
		logs.Error(info)
		logs.Emergency(info)
		logs.Notice(info)
	}
	

	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}