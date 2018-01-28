package main

import (
	"fmt"
	"time"
)

/**
非缓冲 channel 示例
通道长度为0,未分配缓冲空间,遵循 happens before 原则:

1. 向此类通道发送元素值的操作会被阻塞,直到至少有一个接收方,该接收操作会先得到元素值的<副本>,
然后唤醒发送方的 goroutinue 之后返回,即这时的接收操作会在对应的发送操作之前完成
2. 此类通道接收元素值的操作同样被阻塞,直到至少有一个发送方,该发送操作会直接把<元素值>复制给接收方
然后唤醒接收方所在的 goroutunue,即这时的发送操作会在对应的接收操作完成之前完成

只有在非缓冲通道的接收方与发送方双方"握手"之后,元素值的传递才会进行,如果有多个发送方/接收方,那么就需要排队"握手"
*/
func main() {
	sendingInterval := time.Second
	receptionInterval := time.Second * 2
	intChan := make(chan int,0)

	go func(){
		var ts0, ts1 int64
		for i := 1; i< 5; i++ {
			intChan <- i
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("send: ", i)
			} else {
				fmt.Printf("sent: %d [interval: %d s] \n",i,ts1 - ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(sendingInterval)
		}
		close(intChan)
	}()

	var ts0, ts1 int64

	LOOP:
	for {
		select {
		case v, ok := <-intChan:
			if !ok {
				break LOOP
			}
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("received: ",v)
			} else {
				fmt.Printf("received: %d [interval: %d s] \n", v, ts1 - ts0)
			}
		}
		ts0 = time.Now().Unix()
		time.Sleep(receptionInterval)
	}

	fmt.Println("end")
}

