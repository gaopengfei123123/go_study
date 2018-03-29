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

	// 完全成功的接口
	r.POST("/api/test/try", trySuccess)
	r.POST("/api/test/confirm", confirmSuccess)
	r.POST("/api/test/cancel", cancelSuccess)

	// try 中存在失败的接口
	r.POST("/api/test_try_false/try", tryFault)
	

	// commit 中存在失败但可以回滚成功的接口
	r.POST("/api/test_confirm_fault/try", trySuccess)
	r.POST("/api/test_confirm_fault/confirm", confirmFault)
	r.POST("/api/test_confirm_fault/cancel", cancelSuccess)


	// 执行 cancel 回滚不成功的接口
	r.POST("/api/test_cancel_fault/try", trySuccess)
	r.POST("/api/test_cancel_fault/confirm", confirmSuccess)
	r.POST("/api/test_cancel_fault/cancel", cancelFault)


	// 执行 cancel 回滚超时的接口
	r.POST("/api/test_cancel_timeout/try", trySuccess)
	r.POST("/api/test_cancel_timeout/confirm", confirmSuccess)
	r.POST("/api/test_cancel_timeout/cancel", cancelTimeout)





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
		testForm.Type = "cancel api"
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

// cancel 超时
func cancelTimeout(c *gin.Context){
	var statusCode int
	var errorCode int
	var errorMsg string
	errorCode = 408
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
		c.JSON(408, testForm)
	} else {
		// 处理失败时的返回
		c.JSON(408, gin.H{
			"error_code": 408,
			"error_message": "time out",
		})
	}
}