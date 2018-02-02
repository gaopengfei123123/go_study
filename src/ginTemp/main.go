package main

import (
	"ginTemp/router"
	"github.com/gin-gonic/gin" 
)

func main() {
	r := gin.Default()
	r = router.RegistRouter(r)

	r.Run(":8899")
}