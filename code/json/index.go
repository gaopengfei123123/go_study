package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type allStruct interface{}

//定义了获取 json 的结构体
type Package struct {
	Mingzi       string `json:"name"`
	Version      string
	PrIvate      bool
	Dependencies Depend
	Browserslist []string
}
type Depend struct {
	Vue    string
	Router string `json:"vue-router"`
}

func main() {
	path := "./package.json"

	pkg := Package{}
	file, err := os.Open(path)
	checkError(err)
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	checkError(err)
	//这个就是解析 json 格式内容的函数
	json.Unmarshal(data, &pkg)
	foreachStruct(pkg)
	//output
	// Mingzi -- hexo-site
	// Version -- 0.0.0
	// PrIvate -- true
	// Dependencies -- {^2.2.6 ^2.3.1}
	// Browserslist -- [> 1% last 2 versions not ie <= 8]
}

//检测错误的方法
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//遍历结构体的方法
func foreachStruct(st interface{}) {
	t := reflect.TypeOf(st)
	v := reflect.ValueOf(st)

	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		value := v.Field(k).Interface()
		fmt.Printf("%s -- %v \n", key, value)
	}
}
