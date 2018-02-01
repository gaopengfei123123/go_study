package controllers
import (
	"github.com/gin-gonic/gin" 
)
func HelloPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hahaha",
	})
}