package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	在典型的 RESTful 应用中，你可能会在任何路由中遇到错误——无效输入、数据库故障、未授权访问或内部 bug。
	在每个处理函数中单独处理错误会导致重复代码和不一致的响应

	集中式的错误处理中间件通过在每个请求后运行并检查通过 c.Error(err) 添加到 Gin 上下文中的任何错误来解决这个问题。
	如果发现错误，它会发送一个带有正确状态码的结构化 JSON 响应
*/

// 错误处理程序会捕获错误并返回一个格式一致的 JSON 错误响应
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 首先处理该请求

		// 检查是否在上下文中添加了任何错误信息
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}
	}
}

func ErrorHandlingMiddleware(c *gin.Context) {
	c.Error(errors.New("something went wrong"))
}
