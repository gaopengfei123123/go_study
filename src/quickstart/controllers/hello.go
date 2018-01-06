package controllers

import (
	"github.com/astaxie/beego"
)

// HelloController : 用于快速学习 beego 框架
type HelloController struct {
	beego.Controller
}

// Get : 用于承接访问请求
func (c *HelloController) Get() {
	name := c.Ctx.Input.Param(":name")
	c.Ctx.WriteString("hello " + name)
}

// Post : 更新资源
func (c *HelloController) Post() {
	name := c.Ctx.Input.Param(":name")
	c.Ctx.WriteString("this is post master:" + name)
}
