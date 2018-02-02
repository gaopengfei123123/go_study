package main

import (
	"ginTemp/router"
	"ginTemp/middleware"
	"github.com/gin-gonic/gin" 
)

func main() {
	r := gin.Default()
	r = middleware.GlobalMiddleware(r)
	r = router.RegistRouter(r)

	r.Run(":8899")
}