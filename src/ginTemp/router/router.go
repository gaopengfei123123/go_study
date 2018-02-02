package router


import (
	"github.com/gin-gonic/gin" 
	"ginTemp/controllers"
	"net/http"
)

// RegistRouter 注册路由
func RegistRouter(r *gin.Engine) *gin.Engine {
	// 指定访问的静态文件
	r.StaticFile("/", "./view/index.html")
	// 指定访问的目录
	r.StaticFS("/static", http.Dir("./view/static"))

	// 一个正常的get请求示例
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 这里引入了外部包的方法
	r.GET("/hello", controllers.HelloPage)
	r.GET("/hello/:name", controllers.HelloParam)
	r.POST("/hello/login", controllers.HelloForm)
	return r
}
