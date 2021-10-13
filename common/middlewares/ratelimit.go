package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// RateLimitMiddleware 令牌桶中间件 (fillInterval time.Duration, cap int64)-->(时间，个数)
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if bucket.TakeAvailable(1) < 1 { // 或者bucket.TakeAvailable(1) == 0
			c.String(http.StatusOK, "请求过多，请稍后再试...")
			c.Abort()
			return
		}
		// 取到令牌放行
		c.Next()
	}
}
