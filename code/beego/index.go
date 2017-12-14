package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("Hello world")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
