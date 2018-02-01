package main

import (
	// "io"
	"ginTemp/controllers"
	// "time"
	"github.com/gin-gonic/gin" 
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello", controllers.HelloPage)

	r.Run(":8899")
}