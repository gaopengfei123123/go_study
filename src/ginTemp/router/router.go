package router


import (
	"github.com/gin-gonic/gin" 
	"ginTemp/controllers"
	"net/http"
)

// RegistRouter 注册路由
func RegistRouter(r *gin.Engine) *gin.Engine {
	// r.Static("/", "./view/index.html")
	r.StaticFile("/", "./view/index.html")
	r.StaticFS("/static", http.Dir("./view/static"))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello", controllers.HelloPage)
	r.GET("/hello/:name", controllers.HelloParam)
	r.POST("/hello/login", controllers.HelloForm)
	return r
}
