package controllers
import (
	"github.com/gin-gonic/gin" 
	// "fmt"
	// "net/http"
)
// HelloPage 基本示例,返回 json 格式内容
func HelloPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hahaha",
	})
}
// HelloParam 读取路由中带来的参数
func HelloParam(c * gin.Context){
	name := c.Param("name")
	c.String(200, "参数为 %s", name)
}

// LoginForm 在标签中就是映射要读取的字段,同时可以进行字段验证
type LoginForm struct {
	User     string `form:"user" json:"user" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

// HelloForm 将form表单提绑定到 struct 上,form-data
// 这里有 Bind, BindJSON, BindQuery 和 ShouldBind, ShouldBindJSON, ShouldBindQuery 这6中绑定方法
// 加上 should 就是根据 content-type 一定的自我判定能力,代价就是损耗一些性能
func HelloForm(c *gin.Context){
	var json LoginForm
	 // binding JSON,本质是将request中的Body中的数据按照JSON格式解析到json变量中
	 if c.Bind(&json) == nil {
     	c.JSON(404, json)
    } else {
        c.JSON(404, gin.H{"JSON=== status": "binding JSON error!"})
	}
	
}