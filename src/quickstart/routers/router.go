package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 所有的 options 请求走的同一个地址
	beego.Router(`*`, &controllers.BaseController{}, "options:Options")
	beego.Router("/", &controllers.MainController{})
	beego.Router(`/hello/:name([\w]+)`, &controllers.HelloController{})
	beego.AutoRouter(&controllers.ObjectController{})
	beego.Include(&controllers.UserController{})
}
