package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bitly/go-simplejson"
)

func main() {
	filePath := "./composer.json"
	file, err := os.Open(filePath)
	checkError(err)
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	// checkError(err)
	readJSON(data)

	js, err := simplejson.NewJson(data)
	checkError(err)

	//获取某个字段值
	str, err := js.Get("name").String()
	checkError(err)
	fmt.Println("name -> ", str)

	//多层级的key值
	str2, err := js.Get("autoload").Get("classmap").GetIndex(0).String()
	checkError(err)
	fmt.Println("autoload.classmap[0] -> ", str2)

	//判断字段是否存在
	//源码内容 https://github.com/bitly/go-simplejson/blob/master/simplejson.go#L157
	jType, ok := js.CheckGet("type")
	if ok {
		str3, _ := jType.String()
		fmt.Println("type -> ", str3)
	} else {
		fmt.Println("no exist")
	}

	//数组
	arr, err := js.Get("keywords").Array()
	checkError(err)
	for i, v := range arr {
		fmt.Printf("arr index:%d value:%s \n", i, v)
	}

	//字典
	mp := js.Get("require").MustMap()
	fmt.Println("require's key:value is:")
	for key, value := range mp {
		fmt.Printf("%s : %s \n", key, value)
	}

}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readJSON(data []byte) {
	//定义一个由空接口组成的字典用来承接解析出来的 json
	var f map[string]interface{}
	json.Unmarshal(data, &f)
	//通过断言来判断字段类型
	for k, v := range f {
		switch v.(type) {
		case string:
			fmt.Println(k, "is string ====>", v)
		case int:
			fmt.Println(k, "is int ====>", v)
		case []interface{}:
			fmt.Println(k, "is interface ====>", v)
		default:
			fmt.Println(k, "is null type ====>", v)
		}
	}
}
