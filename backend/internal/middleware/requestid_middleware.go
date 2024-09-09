package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

// RequestIDMiddleware 是自定义中间件，用于生成trace_id并注入到上下文中
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成一个唯一的trace_id
		traceID := uuid.New().String()

		// 将trace_id添加到请求头中，以便在响应中传递给客户端
		c.Header("X-Trace-ID", traceID)

		// 将trace_id添加到上下文中，以便后续处理程序可以访问它
		c.Set("trace_id", traceID)

		method := c.Request.Method
		log.Printf("request method: %s\n", method)

		// 根据请求方式，记录header和参数

		// 打印请求日志，包括trace_id
		log.Printf("Incoming Request - trace_id: %s", traceID)

		// 继续处理请求
		c.Next()

		// 打印响应日志，包括trace_id
		log.Printf("Outgoing Response - trace_id: %s", traceID)
	}
}
