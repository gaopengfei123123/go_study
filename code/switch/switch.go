package main

import (
	"fmt"
	"strconv"
)

type Element interface{} //定义一个空接口
type List []Element      //定义一个由空接口组成的数组类型

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "< " + p.name + "-" + strconv.Itoa(p.age) + "years old />"
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "shakalaka"
	list[2] = Person{"Gao", 25}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is int and value is %d \n", index, value)
		case string:
			fmt.Printf("list[%d] is string and value is %s \n", index, value)
		case Person:
			fmt.Printf("list[%d] is Person and value is %s \n", index, value)
		default:
			fmt.Printf("list[%d] is undefined", index)
		}
	}
}
