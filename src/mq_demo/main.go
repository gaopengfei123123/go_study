package main

import(
	"fmt"
	mq "mq_demo/mqhandler"
	"time"
)

func main(){
	fmt.Println("hello mq_handler")
	var handler = mq.MQService{}
	go sendSomething(handler)
	handler.Read(tt)
}

func tt(jsonStr []byte){
	time.Sleep(time.Second * 1)
	fmt.Println("already in listenning mq: ", string(jsonStr))
}


func sendSomething(handler mq.MQ){
	time.Sleep(time.Second * 1)
	handler.Send("testKey1", "testValue1")

	time.Sleep(time.Second * 1)
	handler.Send("testKey2", "testValue2")

	time.Sleep(time.Second * 1)
	handler.Send("testKey3", "testValue3")
}