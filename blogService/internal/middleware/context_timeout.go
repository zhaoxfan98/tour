package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		//将设置了超时的ctx传入
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

/*
如果在进行多应用/服务的调用时，把父级的上下文信息不断地传递下去，那么在统计超时控制的中间件中所设置的超时时间，
其实是针对整条链路的。如果需要单独调整某条链路的超时时间，那么只需调用context.WithTimeout等方法对父级ctx进行设置，
然后取得子级ctx，再进行新的传递即可。
*/
