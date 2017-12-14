package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	handleSignal()
}

//演示接收信号的示例函数
func handleSignal() {
	fmt.Println("hello go")

	// 创建一个信号通道
	sigRecv1 := make(chan os.Signal, 1)
	// 只自行处理 SIGINT 和 SIGQUIT 的信号
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("set notification for %s... [sigRecv1]\n", sigs1)
	//当收到信号时将信号发送给相应的信号值中(因为集合中的值已经满足了 os.Signal 的接口),下同
	signal.Notify(sigRecv1, sigs1...)

	// 创建另一个信号通道
	sigRecv2 := make(chan os.Signal, 1)
	// 只自行处理 SIGQUIT 的信号与集合一形成对比
	sigs2 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("set notification for %s... [sigRecv2]\n", sigs2)
	signal.Notify(sigRecv2, sigs2...)

	// 开启同步执行
	var wg sync.WaitGroup
	//预先设置2个差量,代表两条线程并行
	wg.Add(2)

	// go 关键字开启线程
	go func() {
		// 显示出收到的信号,在信号通道关闭之前会一直阻塞着
		for sig := range sigRecv1 {
			fmt.Printf("Receive a signal from sigRecv1: %s\n", sig)
		}
		fmt.Printf("End. [sigRecv1]\n")
		// 差值 -1
		wg.Done()
	}()
	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("Receive a signal from sigRecv2: %s\n", sig)
		}
		fmt.Printf("End. [sigRecv2]\n")
		wg.Done()
	}()

	fmt.Println("Wait for 2 seconds...")
	// 这里延迟10秒中方便进行信号操作
	time.Sleep(10 * time.Second)
	fmt.Printf("Stop notification....")
	// 停止接收信号(这一步很重要,如果不关闭的话 ctrl+c 就别想正常退出了,只能走 kill -9 的路线)
	// 只关闭通道1是因为只有通道1有自行处理 SIGINT 的信号,通道2没处理过,也就不会发生信号不能执行的问题
	signal.Stop(sigRecv1)
	close(sigRecv1)
	fmt.Printf("done. [sigRecv1]\n")
	// 避免示例函数提前退出(线程被回收)
	wg.Wait()

	// 敲完以上代码后,运行一下该文件, 然后键盘 ctrl + c (SIGINT) 和 ctrl + \ (SIGQUIT) 体验一下
}
