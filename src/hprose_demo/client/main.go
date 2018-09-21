package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	// "time"
)

type HelloService struct {
	Hello    func(string) (string, error)
	CallBack func(module, class, method string, args interface{}) (interface{}, error) `name:"baseServer_callBack"`
}

type BaseService struct {
	Callback func(module, class, method string, args map[string]interface{}) (interface{}, error) `name:"baseServer_callBack"`
}

func main() {
	client := rpc.NewClient("tcp://0.0.0.0:4321")
	var hello *HelloService
	client.UseService(&hello)
	fmt.Println(hello)
	res, err := hello.Hello("from client")
	fmt.Println(res, err)

	// var base *BaseService
	// client.UseService(&base)
	// client.SetTimeout(time.Second * 5)
	// fmt.Println(base)
	// fmt.Println(client.URI())
	// args := map[string]interface{}{
	// 	"mobile": "18333636949",
	// }
	// res, err := base.Callback("System", "XyToken", "getXyOpenKey", args)
	// fmt.Println(res, err)
}
