package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/models"
	"regexp"
)

// UserController 有关于 User 的各种秦操作
type UserController struct {
	beego.Controller
}

type Any interface{}

type responseBody struct {
	Code int	`json:"code"`
	Data Any	`json:"data"`
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
	user := models.UserModal{}
	user.Id = 1
	data := user.GetOne()
	c.Data["json"] = data
	c.ServeJSON()
}


type loginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
// @router /user/login [post]
func (c *UserController) Login() {
	form := loginForm{}
	var resp responseBody

	resp.Code = 200
	resp.Data = "nothing"

	if err := c.ParseForm(&form); err != nil {
		resp.Code = 401
		resp.Data = "error format"
	} else {
		resp.Data = form
	}

	// check form format
	var errorList []string
	if m, _ := regexp.MatchString("^[a-zA-z0-9!#$]{6,20}$",form.Username); !m{
		errorList = append(errorList, "username format error!")
	}
	if m, _ := regexp.MatchString("^[a-zA-z0-9!#$]{6,20}$",form.Password); !m{
		errorList = append(errorList, "password format error!")
	}
	
	// check account
	user := models.UserModal{}
	user.Name = form.Username
	data := user.GetOne()
	if data.Error != "" {
		errorList = append(errorList,data.Error)
	}
	resp.Data = data

	if len(errorList) > 0 {
		resp.Code = 400
		resp.Data = errorList
	} 

	c.Data["json"] = resp
	c.ServeJSON()
}