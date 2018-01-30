package main

import (
	"time"
	"fmt"
)
//select 的规则:
/**
1. 每个case都必须是一个通信
2. 所有channel表达式都会被求值
3. 所有被发送的表达式都会被求值
4. 如果任意某个通信可以进行，它就执行；其他被忽略。
5. 如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。 
6. 否则：如果有default子句，则执行该语句。
7. 如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
*/
func add(ch chan<- int) {
	for i :=0; i < 10; i++ {
		ch <- i
	}
}

// 因为 select 的特性会导致 time.After 会在 ch 清空之后才执行,这不符合预想的逻辑
func test() {
	ch := make(chan int,10)

	go add(ch)
	LOOP:
	for {
		select {
		case <- time.After(2 * time.Second):
			fmt.Println("timeout")
			break LOOP
		case  n := <- ch :
			fmt.Println(n)
			fmt.Println("sleep one seconds ...")
			time.Sleep(1 * time.Second)
			fmt.Println("sleep one seconds end...")
		}
	}
}


// 将 time.After 换成 time.Ticker 会执行倒计时,但是当两个 case 同时生效的时候,会随机选择一个执行,这也有问题
func test2() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	ch := make(chan int,10)

	go add(ch)
	LOOP:
	for {
		select {
		case <- ticker.C:
			fmt.Println("timeout")
			break LOOP
		case  n := <- ch :
			fmt.Println(n)
			fmt.Println("sleep one seconds ...")
			time.Sleep(1 * time.Second)
			fmt.Println("sleep one seconds end...")
		}
	}
}

// 最终的解决办法就是将倒计时的通道单独拎出来, 因为 select 在同一时间只能执行一条语句, 而且 select 需要加 default,不然会因为 ch 为空时发生阻塞
func test3(){
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	ch := make(chan int,10)

	go add(ch)
	LOOP:
	for {
		select {
		case  n := <- ch :
			fmt.Println(n)
			fmt.Println("sleep one seconds ...")
			time.Sleep(1 * time.Second)
			fmt.Println("sleep one seconds end...")
		default:
		}

		select {
		case <- ticker.C:
			fmt.Println("timeout")
			break LOOP
		default:
		}
	}
}


func main() {
	test3()
}

