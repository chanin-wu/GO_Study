package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 禁用控制台颜色，当将日志写入文件时，您无需使用控制台颜色。
	gin.DisableConsoleColor()

	// 将日志记录到文件中。
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果您需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	gin.DefaultWriter = &lumberjack.Logger{
		Filename:   "gin.log",
		MaxSize:    100, // 兆字节
		MaxBackups: 3,
		MaxAge:     28, // days
	}

	r := gin.Default()

	logging := r.Group("/logging")

	{
		logging.GET("/writeLog", WriteLog)
	}

	r.Run(":8080")
}
