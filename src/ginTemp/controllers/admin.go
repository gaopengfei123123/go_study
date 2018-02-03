package controllers
import(
	"github.com/gin-gonic/gin" 
)

// AdminHello 分组路由
func AdminHello(c *gin.Context) {
	example := c.MustGet("oldman").(string)
	c.JSON(200, gin.H{
		"message": "admin hello",
		"param": example,
	})
}

// AdminHi 分组路由2
func AdminHi(c *gin.Context){
	example := c.MustGet("oldman").(string)
	c.JSON(200, gin.H{
		"message": "admin hi",
		"param": example,
	})
}