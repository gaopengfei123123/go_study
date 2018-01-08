package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["quickstart/controllers:NewController"] = append(beego.GlobalControllerRouter["quickstart/controllers:NewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/new`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["quickstart/controllers:UserController"] = append(beego.GlobalControllerRouter["quickstart/controllers:UserController"],
		beego.ControllerComments{
			Method: "Info",
			Router: `/user/info`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["quickstart/controllers:UserController"] = append(beego.GlobalControllerRouter["quickstart/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/user/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
