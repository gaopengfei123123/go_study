package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/models"
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
	Token string		`json:"token"`
}
// Info 获取用户的 头像,名称,权限
// @router /user/info [post]
func (c *UserController) Info() {
	// roles := []string{"admin"}
	// token := "123sdfqs"
	// info := UserInfo{"http://blog-image.onlyoneip.com/6f333b29.jpg","GPF",roles,token}
	user := models.User{Id:1}
	data := user.GetOne()
	c.Data["json"] = data
	c.ServeJSON()
}