package middlewares

import (
	"GoWeb_Template_CLD/common"
	"GoWeb_Template_CLD/common/jwt"
	"GoWeb_Template_CLD/controller"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 格式：Authorization：Bearer xxx.xxx.xxx
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, common.CodeNeedLogin)
			c.Abort() // 中断当前流程处理程序
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, common.CodeInvalidToken)
			c.Abort() // 中断当前流程处理程序
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, common.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(controller.ContextUserIDKey, mc.UserID)

		c.Next() // 后续的处理请求函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息
	}
}