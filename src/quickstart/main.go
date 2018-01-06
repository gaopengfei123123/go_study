package main

import (
	_ "quickstart/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/down1", "down1")
	beego.Run()
}
