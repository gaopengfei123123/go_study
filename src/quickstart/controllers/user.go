package controllers

import (
	"github.com/astaxie/beego"
)

// UserController 有关于 User 的各种秦操作
type UserController struct {
	beego.Controller
}


// UserInfo 返回的 json 结构体
type UserInfo struct{
	Avatar string		`json:"avatar"`
	Name string			`json:"name"`
	Roles []string		`json:"roles"`
}
// Info 获取用户的 头像,名称,权限
// @router /user/info [get]
func (c *UserController) Info() {
	roles := []string{"admin"}
	info := UserInfo{"http://blog-image.onlyoneip.com/6f333b29.jpg","GPF",roles}
	c.Data["json"] = info
	c.ServeJSON()
}