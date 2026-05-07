package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	如何写入日志文件
		将日志写入文件对于生产应用至关重要，你需要保留请求历史记录以用于调试、审计或监控。
		默认情况下，Gin 将所有日志输出写入 os.Stdout。
		你可以在创建路由器之前通过设置 gin.DefaultWriter 来重定向。

	同时写入文件和控制台
		Go 标准库中的 io.MultiWriter 函数接受多个 io.Writer 值，并将写入复制到所有目标。这在开发时很有用，你既想在终端看到日志，又想将其持久化到磁盘

	生产环境中的日志轮转
		使用 os.Create，每次应用启动时都会截断日志文件。在生产环境中，你通常希望追加到现有日志并根据大小或时间轮转文件。考虑使用日志轮转库如 lumberjack
*/

func WriteLog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "WriteLog",
	})
}
