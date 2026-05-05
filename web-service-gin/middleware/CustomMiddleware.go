package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

/*
	Gin 中间件是一个返回 gin.HandlerFunc 的函数。中间件在主处理函数之前和/或之后运行，这使其适用于日志记录、认证、错误处理和其他横切关注点。

	中间件执行流程
	中间件函数分为两个阶段，由 c.Next() 调用分隔
		c.Next() 之前 — 此处的代码在请求到达主处理函数之前运行。用于设置任务，如记录开始时间、验证令牌或使用 c.Set() 设置上下文值
		c.Next() — 调用链中的下一个处理函数（可能是另一个中间件或最终的路由处理函数）。执行在此暂停，直到所有下游处理函数完成
		c.Next() 之后 — 此处的代码在主处理函数完成后运行。用于清理、记录响应状态或测量延迟

	如果你想完全停止链（例如认证失败时），调用 c.Abort() 而不是 c.Next()。这会阻止任何剩余的处理函数执行。你可以将其与响应结合使用，例如 c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
*/

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置示例变量
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func CustomMiddleware(c *gin.Context) {
	example := c.MustGet("example").(string)

	// it would print: "12345"
	log.Println(example)
}
