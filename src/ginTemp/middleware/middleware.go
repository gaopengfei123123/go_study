package middleware

import (
	"github.com/gin-gonic/gin" 
	"time"
	"log"
)

// Logger 测试中间件
// 这里的 c.Next() 就是controller 中的方法,相当于实际的业务逻辑, c.Next()的位置
// 这里的属于全局的中间件, 这个算是后置中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "233333")
		c.Next()

		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()
		log.Println(status)
		log.Println("this is Logger after func")
	}
}

// Customer2 另一个中间件, 算是函数的前置中间件
func Customer2() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("this is Customer2 before func")
		c.Next()
	}
}

// GlobalMiddleware 注册全局中间件
func GlobalMiddleware(r *gin.Engine) *gin.Engine {
	// 绑定全局中间件,可以注册多个
	r.Use(Logger(),Customer2())

	return r
}