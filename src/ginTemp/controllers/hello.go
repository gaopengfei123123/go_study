package controllers
import (
	"github.com/gin-gonic/gin" 
	// "fmt"
	// "net/http"
)
func HelloPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hahaha",
	})
}

func HelloParam(c * gin.Context){
	name := c.Param("name")
	c.String(200, "参数为 %s", name)
}

type LoginForm struct {
	User     string `form:"user" json:"user" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

// HelloJson 将form表单提绑定到 struct 上,form-data
func HelloForm(c *gin.Context){
	var json LoginForm
	 // binding JSON,本质是将request中的Body中的数据按照JSON格式解析到json变量中
	 if c.Bind(&json) == nil {
     	c.JSON(404, json)
    } else {
        c.JSON(404, gin.H{"JSON=== status": "binding JSON error!"})
	}
	
}