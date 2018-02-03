package router


import (
	"github.com/gin-gonic/gin" 
	"ginTemp/controllers"
	"net/http"
	// "time"
	"log"
)

// func Logger() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		t := time.Now()
// 		c.Set("example", "233333")
// 		c.Next()

// 		latency := time.Since(t)
// 		log.Print(latency)

// 		status := c.Writer.Status()
// 		log.Println(status)
// 	}
// }


// RegistRouter 注册路由
func RegistRouter(r *gin.Engine) *gin.Engine {
	// r.Use(Logger())
	r.POST("/logger", func(c *gin.Context){
		example := c.MustGet("example").(string)
		log.Println("this is loggerFunc")
		c.JSON(200,gin.H{
			"output": example,
		})
	})
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
	// 带参数的路由,以及方法中怎么读取
	r.GET("/hello/:name", controllers.HelloParam)
	// 进行表单提交的时候,该怎么绑定提交的参数
	r.POST("/hello/login", controllers.HelloForm)
	return r
}
