package main

import (
	"fmt"
	"net/http"
	"time"
	"sync"
)

var urls = []string{
	"http://www.baidu.com/",
	"http://www.taobao.com/",
	"http://blog.justwe.site/",
}

func main() {
	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go requestURL(url,&wg)
	}

	wg.Wait()
	fmt.Println("down")
}

// 判断url 的状态
func requestURL(url string ,wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	resp,err := http.Head(url)
	if err != nil {
		fmt.Println("Error:",url, err)
	}
	end := time.Now()
	exeTime := end.Sub(start).Nanoseconds() / 1000000
	fmt.Println(url, ":", resp.Status, " delay:" ,  exeTime , "ms")
}