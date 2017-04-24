package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

// 这里需要把 int 类型转换成字符串,否则使用 + 来连接字符的时候报数据类型不匹配的错误
func (h Human) String() string {
	return "<" + h.name + "-" + strconv.Itoa(h.age) + "years old ,phone is " + h.phone + "/>"
}

func main() {
	Duck := Human{"duck", 22, "110-120-119"}
	fmt.Println("this human is :", Duck)
}
