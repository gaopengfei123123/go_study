package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func main() {
    ch := make(chan int)
    //定义一个WaitGroup，阻塞主线程执行
    var wg sync.WaitGroup
    //添加一个goroutine等待
    wg.Add(1)
    //goroutine超时
    go func() {
        //执行完成，减少一个goroutine等待
        defer wg.Done()
        for {
            select {
            case i := <-ch:
                fmt.Println(i)
            //goroutine内部3秒超时
            case <-time.After(3 * time.Second):
                fmt.Println("goroutine1 timed out")
                return
            }
        }
    }()
    ch <- 1
    //新增一个1秒执行一次的计时器
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    //新增一个10秒超时的上下文
    background := context.Background()
    // 设置一个10秒倒计时的context,过期后默认执行
    ctx, _ := context.WithTimeout(background, 10*time.Second)
    //添加一个goroutine等待
    wg.Add(1)
    go func(ctx context.Context) {
        //执行完成，减少一个goroutine等待
        defer wg.Done()
        for {
            select {
            //每秒一次
            case <-ticker.C:
                fmt.Println("tick")
            //内部超时，不会被执行
            case <-time.After(5 * time.Second):
                fmt.Println("goroutine2 timed out")
            //上下文传递超时信息，结束goroutine
            case <-ctx.Done():
                fmt.Println("goroutine2 done")
                return
            }
        }
    }(ctx)
    //等待所有goroutine执行完成
    wg.Wait()
}
