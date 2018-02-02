package middleware

import (
	"github.com/gin-gonic/gin" 
	"time"
	"log"
)

// Logger 测试中间件
// 这里的 c.Next() 就是controller 中的方法,相当于实际的业务逻辑, c.Next()的位置
// 这里的属于全局的中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "233333")
		c.Next()

		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
}

// GlobalMiddleware 注册全局中间件
func GlobalMiddleware(r *gin.Engine) *gin.Engine {
	r.Use(Logger())

	return r
}