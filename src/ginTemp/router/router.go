package router


import (
	"github.com/gin-gonic/gin" 
	"ginTemp/controllers"
	"ginTemp/middleware"
	"net/http"
	// "time"
	"log"
	"github.com/gin-gonic/gin/binding"
	"reflect"
	"gopkg.in/go-playground/validator.v8"
	"time"
)


// 自定义的验证规则 提交的时间不得晚于当前日期
func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	log.Println("validate time")
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false	// 验证不通过
		}
	}
	return true
}


// RegistRouter 注册路由
func RegistRouter(r *gin.Engine) *gin.Engine {

	binding.Validator.RegisterValidation("bookabledate", bookableDate)
	
	// 指定访问的静态文件
	r.StaticFile("/", "./view/index.html")
	// 指定访问的目录
	r.StaticFS("/static", http.Dir("./view/static"))

	// 一个正常的get请求示例
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 这里引入了外部包的方法
	r.GET("/hello", controllers.HelloPage)
	// 带参数的路由,以及方法中怎么读取
	r.GET("/hello/:name", controllers.HelloParam)
	// 进行表单提交的时候,该怎么绑定提交的参数
	r.POST("/hello/login", controllers.HelloForm)
	
	r.POST("/hello/validate", controllers.HelloValidate)



	// 这里用来演示全局中间件的测试
	r.POST("/logger", func(c *gin.Context){
		example := c.MustGet("example").(string)
		oldman, isExit := c.Get("oldman")
		if !isExit {
			oldman = "no param"
		}
		log.Println("this is loggerFunc")
		c.JSON(200,gin.H{
			"output": example,
			"param": oldman,
		})
	})

	// 用于测试单个中间件绑定路由的例子
	r.POST("/test", middleware.ForTest,controllers.AdminHello)

	// 这里演示了路由分组以及针对分组怎么对分组使用中间件
	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.ForAdmin)
	{
		adminGroup.POST("/hello", controllers.AdminHello)
		adminGroup.POST("/hi", controllers.AdminHi)
	}
	return r
}
