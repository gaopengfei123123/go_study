package controllers

import (
	"github.com/gin-gonic/gin"
	"rpc_ts/services"
)




// Client 串行事务的接口
func Client(c *gin.Context) {
	var request services.ClientForm

	// 验证是否成功绑定
	if c.BindJSON(&request) == nil {
		
		// 真正处理逻辑
		response := services.ClientService(request)
		c.JSON(200, response)
	} else {

		// 处理失败时的返回
		c.JSON(200, gin.H{
			"message" : "get nothing",
		})
	}
}