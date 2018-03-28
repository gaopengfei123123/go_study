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

	r.POST("/api/test/try", func(c *gin.Context){
		var testForm struct{
			Message string `json:"message"`
			Type	string	`json:"type"`
			Code	int		`json:"code"`
			ErrorCode int	`json:"error_code"`
			ErrorMessage string	`json:"error_message"`
			Data string `json:"data"`
		}
		// 验证是否成功绑定
		if c.BindJSON(&testForm) == nil {
			testForm.Type = "try api"
			testForm.Code = 200
			testForm.ErrorCode = 0
			testForm.ErrorMessage = "nothing wrong"
			testForm.Data = "{\"message\": \"try\"}"
			c.JSON(200, testForm)
		} else {
			// 处理失败时的返回
			c.JSON(400, gin.H{
				"error_code": 401,
				"error_message": "get nothing",
			})
		}
	})

	r.POST("/api/test/confirm", func(c *gin.Context){
		var testForm struct{
			Message string `json:"message"`
			Type	string
		}
		// 验证是否成功绑定
		if c.BindJSON(&testForm) == nil {
			testForm.Type = "confirm api"
			c.JSON(200, testForm)
		} else {
			// 处理失败时的返回
			c.JSON(400, gin.H{
				"message" : "get nothing",
			})
		}
	})

	r.POST("/api/test/cancel", func(c *gin.Context){
		var testForm struct{
			Message string `json:"message"`
			Type	string
		}
		// 验证是否成功绑定
		if c.BindJSON(&testForm) == nil {
			testForm.Type = "cancel api"
			c.JSON(200, testForm)
		} else {
			// 处理失败时的返回
			c.JSON(400, gin.H{
				"message" : "get nothing",
			})
		}
	})



	return r
}