package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //将 form 内容解析,不然不能在后端显示出来

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprint(w, "welcome!")
}

func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method) //获取提交的方式
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl") //渲染指定的模板
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		//从打印出来的结果来看它不是一个字符串,而是一个 slice,因此如果有多个 username 将
		//存放在一个切片当中,就算只提交一个 username 也是放在这个 slice 中,这点要注意
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
