package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router(`/hello/:name([\w]+)`, &controllers.HelloController{})
	beego.AutoRouter(&controllers.ObjectController{})
	beego.Include(&controllers.UserController{})
}
