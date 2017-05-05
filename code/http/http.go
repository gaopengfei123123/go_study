package main

import (
	"fmt"
	"net/http"
)

//设定一个空的接口体来承载接口
type MyMux struct{}

//实现 Handler 接口,在这里处理路由相关的内容
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w, r)
		return
	}

	if r.URL.Path == "/login" {
		loginPage(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello route '/' !")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is login page")
}

func main() {
	//实例化结构体
	mux := &MyMux{}
	//当做 handler 注入到监听服务当中去
	http.ListenAndServe(":9090", mux)
}
