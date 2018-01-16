package controllers

import (
	// "github.com/astaxie/beego"
	"quickstart/models"
	"regexp"
	"fmt"
)

// UserController 有关于 User 的各种秦操作
type UserController struct {
	BaseController
}

// Any 定义一个空接口
type Any interface{}

// ResponseBody 数据的返回格式
type ResponseBody struct {
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
	c.AllowCross()
	token := c.GetString("token")
	
	user := models.UserModal{}
	user.Token =token
	data := user.GetOne("Token")
	data.GetRoles()

	if data.IsError {
		returnError(c,400,data.Error)
		return
	}

	var userInfo UserInfo
	userInfo.Name = data.User.Name
	userInfo.Avatar = data.User.Avatar
	var roles []string
	for _,item := range data.Roles {
		roles = append(roles,item.Role)
	}
	userInfo.Roles = roles
	userInfo.Token = data.User.Token

	fmt.Println("log:",userInfo)


	resp := ResponseBody{Code:200,Data: userInfo}
	c.Data["json"] = resp
	c.ServeJSON()
}


type loginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// Login 登录操作
// @router /user/login [post]
func (c *UserController) Login() {
	c.AllowCross()
	form := loginForm{}
	var resp ResponseBody

	resp.Code = 200
	resp.Data = "nothing"

	if err := c.ParseForm(&form); err != nil {
		resp.Code = 401
		resp.Data = "error format"
	} else {
		resp.Data = form
	}

	// check form format
	errorList := form.checkLoginForm()
	if len(errorList) > 0 {
		returnError(c,400,errorList)
		return
	} 
	
	// check account
	user := models.UserModal{}
	user.Name = form.Username
	user.Password = form.Password
	user.Ip = c.Ctx.Input.IP()

	data := user.Login()
	if data.IsError {
		returnError(c,400,data.Error)
		return
	}
	user.GetRoles()
	var roles []string
	if len(user.Roles) > 0 {
		for _,item := range data.Roles {
			roles = append(roles,item.Role)
		}
	}else{
		roles = append(roles,"default")
	}
	

	response := struct{
		Token string `json:"token"`
		Role []string `json:"role"`
		Avatar string `json:"avatar"` 
	}{user.Token,roles,user.Avatar}
	resp.Data = response

	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *loginForm) checkLoginForm() []string {
	var errorList []string
	if m, _ := regexp.MatchString("^[a-zA-z0-9!#$]{6,20}$",c.Username); !m{
		errorList = append(errorList, "username format error!")
	}
	if m, _ := regexp.MatchString("^[a-zA-z0-9!#$]{6,20}$",c.Password); !m{
		errorList = append(errorList, "password format error!")
	}
	return errorList
}

func returnError(c *UserController,code int,data Any){
	body := ResponseBody{Code: code,Data: data}
	c.Data["json"] = body
	c.ServeJSON()
}