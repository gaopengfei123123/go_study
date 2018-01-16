package controllers
import (
	"github.com/astaxie/beego"
)
// BaseController 所有类继承的基础类
type BaseController struct {
	beego.Controller
}

// AllowCross 运行跨域
func (c *BaseController) AllowCross() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")       //允许访问源
    c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")    //允许post访问
    c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Token") //header的类型
    c.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
    c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
    c.Ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
}

// Options 面对复杂请求的时候出现的 option method 请求
func (c *BaseController) Options() {
	c.AllowCross()
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}