package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Gin 框架中重定向的用途
URL 结构调整：假设你的网站重构了，原来的一些 URL 不再适用。比如你把某个功能从 /old - feature 迁移到了 /new - feature，为了保证老的链接仍然可用，就可以设置从 /old - feature 重定向到 /new - feature。这样，即使访问老链接，用户也能正常访问到新功能。
用户登录与导航：当用户成功登录后，你可能希望将用户重定向到用户的个人主页。例如，用户在 /login 页面登录成功后，重定向到 /user/dashboard，这样用户就能直接进入他们的操作界面。
资源移动：如果某些资源（如图片、文档等）的存储位置发生了变化，通过重定向可以让访问旧位置的请求自动转到新位置，确保资源的正常访问。
*/
func Redirects() {
	router := gin.Default()

	// External redirect (GET) 外部重定向（GET）
	router.GET("/old", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	// Redirect from POST -- use 302 or 307 to preserve behavior
	// 从 POST 方式重定向——使用 302 或 307 状态码以保持原有行为
	router.POST("/submit", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/result")
	})

	// Internal router redirect (no HTTP round-trip)
	// 内部路由器重定向（无 HTTP 往返过程）
	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/final"
		router.HandleContext(c)
	})

	router.GET("/final", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	router.GET("/result", func(c *gin.Context) {
		c.String(http.StatusOK, "Redirected here!")
	})

	router.Run(":8080")
}
