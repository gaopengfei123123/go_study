package main

import(
  "fmt"
  "net/http"
  "log"
  "html/template"
)


func sayHello(w http.ResponseWriter, r *http.Request){
  t, _ := template.ParseFiles("index.gtpl") //渲染指定的模板
  t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method) //获取提交的方式
	if r.Method == "POST" {
    r.ParseForm()
		//从打印出来的结果来看它不是一个字符串,而是一个 slice,因此如果有多个 username 将
		//存放在一个切片当中,就算只提交一个 username 也是放在这个 slice 中,这点要注意
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
    fmt.Println("are u ok:", r.Form["is_ok"])
    // fmt.Fprint(w, r.Form)

    //输出转以后的字符串
    es_html := template.HTMLEscapeString(r.Form.Get("username"))
    //输出到服务器端
    fmt.Println("username:", es_html)
    //转义并写入到 response 输出到客户端
    template.HTMLEscape(w, []byte(r.Form.Get("username")))
    //原样输出 html 标签
    test := template.HTML(r.Form.Get("username"));
    //返回给客户端
    w.Write([]byte(test))
	}
}


func main(){
  fmt.Println("server is started")

  http.HandleFunc("/", sayHello)
  http.HandleFunc("/login", login)

  error := http.ListenAndServe(":9090", nil)

  if error != nil{
    log.Fatal("ListenAndServe:",error)
  }
}
