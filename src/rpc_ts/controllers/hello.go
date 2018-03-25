package controllers

import (
	"github.com/gin-gonic/gin"
)


// HelloPage 基本示例,返回 json 格式内容
func HelloPage (c *gin.Context) {
	c.JSON(200, gin.H{
		"message" : "hello world",
	})
}