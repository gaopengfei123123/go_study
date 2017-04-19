package main

import (
  "fmt"
  "net/http"
  "strings"
  "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()                         //解析参数,默认不解析
  fmt.Println(r.Form)                   //在终端中打印出表单内容
  fmt.Println("Path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])       //输出指定的参数
  for k,v := range r.Form{              //遍历打印出表单的值
    fmt.Println("key",k)
    fmt.Println("val", strings.Join(v, ""))
  }
  fmt.Fprint(w, "hello gpf!")
}


func main(){
  http.HandleFunc("/", sayhelloName)        //绑定路由与方法
  err := http.ListenAndServe(":9090", nil)  //监听 tcp:9090 端口
  if err != nil {
    log.Fatal("ListenAndServe: ",err)
  }
}
