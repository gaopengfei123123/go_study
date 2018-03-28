package router
import (
	"github.com/gin-gonic/gin"
	"rpc_ts/controllers"
	"net/http"
	"time"
)


// RegistRouter 注册路由
func RegistRouter(r *gin.Engine) *gin.Engine {
	// 指定访问的静态文件
	r.StaticFile("/", "./view/index.html")
	// 指定访问的目录
	r.StaticFS("/static", http.Dir("./view/static"))

	r.GET("/hello", controllers.HelloPage)

	// 接收执行事务的接口  Serial(串行)  Parallel(并行)
	r.POST("/client", controllers.Client)

	r.POST("/api/test", func(c *gin.Context){
		var testForm struct{
			Message string `json:"message"`
		}

		time.Sleep(time.Second * 2)

		// 验证是否成功绑定
		if c.BindJSON(&testForm) == nil {
			
			c.JSON(200, testForm)
		} else {
			// 处理失败时的返回
			c.JSON(200, gin.H{
				"message" : "get nothing",
			})
		}

	})



	return r
}