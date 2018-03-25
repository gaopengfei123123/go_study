package router
import (
	"github.com/gin-gonic/gin"
	"rpc_ts/controllers"
	"net/http"
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



	return r
}