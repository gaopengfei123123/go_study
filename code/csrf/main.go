package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
)

type httpServer struct {
}

func (s *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Stop here if its Preflighted OPTIONS request
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")	// 限制请求方的域 * 就是不限制
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")	// 限制请求方式
		w.Header().Set("Access-Control-Allow-Headers","Action, Module")   //有使用自定义头 需要这个,Action, Module是例子
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Token") //允许的header的类型
	}

	// if r.Method == "OPTIONS" {	// 针对复杂请求时做出的处理
	// 	return
	// }

	w.Write([]byte("hello"))
}

func main() {
	addr := flag.String("http-address", "", "")
	flag.Parse()

	var h httpServer

	httpListener, err := net.Listen("tcp", *addr)
	server := http.Server{Handler: &h,}
	server.Serve(httpListener)

	fmt.Println("finish ", err)
}