package main

import (
	"github.com/gin-gonic/gin"
	"rpc_ts/router"
	"rpc_ts/services"
	// "sync"
)

func main() {
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	var mq services.MQService
	go mq.Read()
	// wg.Wait()

	r := gin.Default()

	r = router.RegistRouter(r)
	r.Run(":8899") // listen and serve on 0.0.0.0:8080
}