package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// 禁用控制台颜色，当将日志写入文件时，您无需使用控制台颜色。
	// gin.DisableConsoleColor() // 禁用日志着色
	gin.ForceConsoleColor() // 启用日志着色

	// 将日志记录到文件中。
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果您需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 追加到现有日志
	/* gin.DefaultWriter = &lumberjack.Logger{
		Filename:   "gin.log",
		MaxSize:    100, // 兆字节
		MaxBackups: 3,
		MaxAge:     28, // days
	} */

	/* ------------------------------------上下分隔------------------------------------ */

	// 通过在“日志配置”中设置“跳过路径”选项，可对指定路径的日志记录进行跳过处理。
	loggerConfig := gin.LoggerConfig{SkipPaths: []string{"/writeLog", "/customLogFormat"}, SkipQueryString: true}

	// 通过在 LoggerConfig 中设置 Skip 函数，您可以根据自己的逻辑跳过日志记录。
	loggerConfig.Skip = func(c *gin.Context) bool {
		// 例如，忽略非服务器端错误
		return c.Writer.Status() < http.StatusInternalServerError
	}

	r.Use(gin.LoggerWithConfig(loggerConfig))

	/* ------------------------------------上下分隔------------------------------------ */

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 您的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// r.Use(gin.Logger())
	r.Use(gin.Recovery())

	/* ------------------------------------上下分隔------------------------------------ */

	// 通过为 gin.DebugPrintRouteFunc 赋值一个函数来自定义此格式
	// 定义路由日志格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	/* ------------------------------------上下分隔------------------------------------ */

	logging := r.Group("/logging")

	{
		logging.GET("/writeLog", WriteLog)
	}

	{
		logging.GET("/customLogFormat", CustomLogFormat)
	}

	{
		logging.GET("/skipLogging", SkipLogging)
	}

	{
		logging.GET("/controllingLogOutputColoring", ControllingLogOutputColoring)
	}

	{
		logging.GET("/avoidLoggingQueryStrings", AvoidLoggingQueryStrings)
	}

	r.Run(":8080")
}
