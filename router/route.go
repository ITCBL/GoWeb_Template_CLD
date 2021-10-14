package router

import (
	"GoWeb_Template_CLD/common/logger"
	"GoWeb_Template_CLD/common/middlewares"
	"GoWeb_Template_CLD/controller"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin 设置成发布模式。发布模式：程序启动时，不将[GIN-debug]信息打印到控制台
	}
	r := gin.New()

	// 限流中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))

	// 静态文件（前端）
	r.LoadHTMLFiles("./templates/index.html") // 导入首页文件
	r.Static("/static", "./static")           // relativePath:为了给前端index.html指定静态的访问路径。 root:静态文件的实际路径
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1 := r.Group("/api/v1/use") // 注册路由组
	var controller controller.UserController

	{ // 不需登陆
		// 注册业务路由
		v1.POST("/register", controller.Register)
		v1.POST("/login", controller.Login)
	}

	{ // 需要登录
		useGroup := v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件
		useGroup.GET("/test", nil)
	}

	// 不存在的请求
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "页面走丢了。。。",
		})
	})

	return r
}
