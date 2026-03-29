package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func loginEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"action": "login"})
}

func submitEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"action": "submit"})
}

func readEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"action": "read"})
}

// 日志记录中间件
func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Printf("请求开始于: %v\n", start)
		c.Next()
		end := time.Now()
		fmt.Printf("请求结束于: %v, 耗时: %v\n", end, end.Sub(start))
	}
}

func GroupingRoutes() {
	router := gin.Default()

	// Simple group: v1
	{
		v1 := router.Group("/v1")
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	{
		v2 := router.Group("/v2")
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	// Public routes -- no auth required
	public := router.Group("/api")
	{
		public.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "健康检查")
		})
	}

	// Private routes -- auth middleware applied to the whole group
	private := router.Group("/api")
	private.Use(loggerMiddleware())
	{
		private.GET("/profile", func(c *gin.Context) {
			c.JSON(http.StatusOK, "获取个人信息")
		})
		private.POST("/settings", func(c *gin.Context) {
			c.JSON(http.StatusOK, "更新设置")
		})
	}

	router.Run(":8080")
}
