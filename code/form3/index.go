package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func heyGuys(w http.ResponseWriter, r *http.Request) {
	//根据时间戳md5加密生成 token
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))

	t, _ := template.ParseFiles("index.gtpl")
	//传入模板当中 {{ . }} 的位置
	t.Execute(w, token)
}

func viewPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form)

		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func main() {
	fmt.Println("starting...")
	http.HandleFunc("/", heyGuys)
	http.HandleFunc("/view", viewPage)

	error := http.ListenAndServe(":9090", nil)

	if error != nil {
		log.Fatal("Whoops, looks like something went wrong")
	}
}
