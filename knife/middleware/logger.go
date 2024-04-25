package middleware

import (
	"knife"
	"time"
)

// Logger 实例一个 Logger 中间件
func Logger() knife.MiddlewareFunc {
	return func(c *knife.Context) {
		// Start timer
		start := time.Now()
		path := c.Req.URL.Path

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.Req.RemoteAddr
		method := c.Req.Method
		statusCode := c.Writer.Status()

		knife.Logger.Printf("remoteAddr:%s method:%s statusCode:%d time:%d path:%s",
			clientIP,
			method,
			statusCode,
			latency,
			path,
		)
	}
}
