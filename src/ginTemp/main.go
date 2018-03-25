package main

import (
	"ginTemp/router"
	"ginTemp/middleware"
	"github.com/gin-gonic/gin" 
)

func main() {
	r := gin.Default()
	// 全局注册的中间件
	r = middleware.GlobalMiddleware(r)
	r = router.RegistRouter(r)

	r.Run(":8899")
}