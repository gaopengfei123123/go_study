package main

import (
	"github.com/gin-gonic/gin"
	"rpc_ts/router"
)

func main() {
	r := gin.Default()

	r = router.RegistRouter(r)
	r.Run(":8899") // listen and serve on 0.0.0.0:8080
}