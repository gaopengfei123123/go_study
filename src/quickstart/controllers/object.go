package controllers

import (
	"github.com/astaxie/beego"
)

// ObjectController : 用于测试自动路由
type ObjectController struct {
	beego.Controller
}

// Login : 登录
func (c *ObjectController) Login() {
	params := c.Ctx.Input.Params()

	output := ""
	for k, v := range params {
		output += " key:" + k + " value:" + v + "; "
	}

	c.Ctx.WriteString("this is login page, params is " + output)
}
