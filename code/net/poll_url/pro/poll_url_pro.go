package main

import (
	"log"
	"net/http"
	"time"
	"sync"
	"context"
)

var urls = []string{
	"http://www.baidu.com/",
	"http://www.taobao.com/",
	"http://blog.justwe.site/",
}

// 加强版的 poll_url , 这次添加了执行超时
func main() {
	pollURL(urls)
}

var executeTimeout = 2 * time.Second

func pollURL(urls []string) {
	log.Printf("Timeout occurs if the execution exceeds %s seconds.\n", executeTimeout)
	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go requestURL2(url, &wg)
	}
	wg.Wait()

	log.Println("all clear")
}

// 判断url 的状态
func requestURL(url string ,wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	// 来自官方 http 包的执行方案,更简单
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp,err := client.Head(url)
	if err != nil {
		log.Println("Error:",url, err)
	}
	end := time.Now()
	exeTime := end.Sub(start).Nanoseconds() / 1000000
	log.Println(url, ":", resp.Status, " delay:" ,  exeTime , "ms")
}

// 判断 url 状态, 自己实现的一个请求链接超时方法,可能会发生某些问题,暂时想不出来
func requestURL2(url string ,wg *sync.WaitGroup) {
	
	// 因为每个url 请求都是独立的请求,因此有个 recover 就很有必要了
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	// 结束时通知完成也是很有必要的吧?
	defer wg.Done()

	// 创建一个非缓冲的 channel,用来接受 http.Head() 执行完成后的通知,完事之后还是要关闭的,不然会发生通道等待
	dst := make(chan struct{})
	defer close(dst)

	// 这里就是执行 url 请求的地方,因为这里会发生阻塞,所以执行部分和倒计时部分要分开执行,goroutinue 是必然的
	go func() {
		start := time.Now()

		resp,err := http.Head(url)
		if err != nil {
			log.Println("Error:",url, err)
		}
		end := time.Now()
		exeTime := end.Sub(start).Nanoseconds() / 1000000
		log.Println(url, ":", resp.Status, " delay:" ,  exeTime , "ms")
		
		dst <- struct{}{}
	}()

	// 这里就是声明一个倒计时
	ctx, cancel := context.WithTimeout(context.Background(),executeTimeout)
	defer cancel()

	// 监听接收消息,要么就是接收超时,要么就是收到东西
	LOOP:
	for {
		select {
		case <-dst:
			break LOOP
		case <-ctx.Done():
			log.Printf("URL: %s has been running too looong! \n",url)
			break LOOP
		}
	}
}



