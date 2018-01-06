package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "blog.justwe.site"
	c.Data["Email"] = "5173180@qq.com"
	c.TplName = "index.tpl"
}
