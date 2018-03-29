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

	r.POST("/api/test/try", trySuccess)
	r.POST("/api/test/confirm", confirmSuccess)
	r.POST("/api/test/cancel", cancelSuccess)



	return r
}

// try成功
func trySuccess(c *gin.Context){
	var statusCode int
	var errorCode int
	var errorMsg string
	statusCode = 200

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
		testForm.Code = statusCode
		testForm.ErrorCode = errorCode
		testForm.ErrorMessage = errorMsg
		testForm.Data = "{\"message\": \"try\"}"
		c.JSON(200, testForm)
	} else {
		// 处理失败时的返回
		c.JSON(400, gin.H{
			"error_code": 401,
			"error_message": "get nothing",
		})
	}
}

// try 失败
func tryFault(c *gin.Context){
	var statusCode int
	var errorCode int
	var errorMsg string
	errorCode = 401
	errorMsg = "something was wrong"

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
		testForm.Code = statusCode
		testForm.ErrorCode = errorCode
		testForm.ErrorMessage = errorMsg
		testForm.Data = "{\"message\": \"try\"}"
		c.JSON(200, testForm)
	} else {
		// 处理失败时的返回
		c.JSON(400, gin.H{
			"error_code": 401,
			"error_message": "get nothing",
		})
	}
}

// confirm 成功
func confirmSuccess(c *gin.Context){
	var statusCode int
	var errorCode int
	var errorMsg string
	statusCode = 200

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
		testForm.Code = statusCode
		testForm.ErrorCode = errorCode
		testForm.ErrorMessage = errorMsg
		testForm.Data = "{\"message\": \"try\"}"
		c.JSON(200, testForm)
	} else {
		// 处理失败时的返回
		c.JSON(400, gin.H{
			"error_code": 401,
			"error_message": "get nothing",
		})
	}
}

// confirm 失败
func confirmFault(c *gin.Context){
	var statusCode int
	var errorCode int
	var errorMsg string
	errorCode = 401
	errorMsg = "something was wrong"

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
		testForm.Code = statusCode
		testForm.ErrorCode = errorCode
		testForm.ErrorMessage = errorMsg
		testForm.Data = "{\"message\": \"try\"}"
		c.JSON(200, testForm)
	} else {
		// 处理失败时的返回
		c.JSON(400, gin.H{
			"error_code": 401,
			"error_message": "get nothing",
		})
	}
}


func cancelSuccess(c *gin.Context){
	var statusCode int
	var errorCode int
	var errorMsg string
	statusCode = 200

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
		testForm.Code = statusCode
		testForm.ErrorCode = errorCode
		testForm.ErrorMessage = errorMsg
		testForm.Data = "{\"message\": \"try\"}"
		c.JSON(200, testForm)
	} else {
		// 处理失败时的返回
		c.JSON(400, gin.H{
			"error_code": 401,
			"error_message": "get nothing",
		})
	}
}


// cancel 失败
func cancelFault(c *gin.Context){
	var statusCode int
	var errorCode int
	var errorMsg string
	errorCode = 401
	errorMsg = "something was wrong"

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
		testForm.Code = statusCode
		testForm.ErrorCode = errorCode
		testForm.ErrorMessage = errorMsg
		testForm.Data = "{\"message\": \"try\"}"
		c.JSON(200, testForm)
	} else {
		// 处理失败时的返回
		c.JSON(400, gin.H{
			"error_code": 401,
			"error_message": "get nothing",
		})
	}
}