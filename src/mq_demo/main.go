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
	fmt.Println("already in listenning mq: ", string(jsonStr))
}


func sendSomething(handler mq.MQ){
	time.Sleep(time.Second * 1)

	handler.Delay("testKey1", "testValue1", "3000")

	time.Sleep(time.Second * 1)
	handler.Delay("testKey2", "testValue2", "4000")

	time.Sleep(time.Second * 1)
	handler.Delay("testKey3", "testValue3", "5000")

	handler.Delay("delayKey1", "this is delay key233", "6000")
}