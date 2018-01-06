package controllers

import (
	"github.com/astaxie/beego"
)

type NewController struct {
	beego.Controller
}

func (c *NewController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// @router /new [get]
func (c *NewController) Get() {
	c.Ctx.WriteString("from comment inject route")
}
