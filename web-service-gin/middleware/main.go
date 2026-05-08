package main

import (
	"net/http"
	"time"

	"github.com/danielkov/gin-helmet/ginhelmet"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	csrf "github.com/utrack/gin-csrf"
)

func main() {
	router := gin.New()

	// gin-helmet 是一个为 Go 语言 Web 框架提供 HTTP 安全中间件的集合
	router.Use(ginhelmet.Default())

	// 下面两个是 全局中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://example.com"},                     // 指定允许访问该资源的源站列表
		AllowMethods:     []string{"GET", "POST"},                             // 定义允许的 HTTP 方法列表
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 指定允许在请求中包含的 HTTP 头字段
		ExposeHeaders:    []string{"Content-Length"},                          // 定义哪些头字段可以被浏览器访问
		AllowCredentials: true,                                                // 设置为 true 表示允许在跨源请求中携带身份凭证（如 cookies、HTTP 认证及客户端 SSL 证书等）
		MaxAge:           12 * time.Hour,                                      // 指定预检请求（OPTIONS 请求）的结果可以被缓存的最长时间 12 * time.Hour 表示预检请求的结果可以被缓存 12 小时。
	}))

	// CSRF
	// 这行代码创建了一个基于 Cookie 的会话存储。cookie.NewStore 函数接受一个字节切片作为密钥，用于加密和解密存储在 Cookie 中的会话数据
	store := cookie.NewStore([]byte("session-secret"))
	// 这里将刚刚创建的会话存储与名为 "mysession" 的会话关联起来，并将这个会话中间件应用到整个 Gin 路由器 router 上。每个请求经过这个中间件时，都会初始化或恢复与 "mysession" 相关的会话
	router.Use(sessions.Sessions("mysession", store))

	router.Use(csrf.Middleware(csrf.Options{
		Secret: "csrf-token-secret",
		ErrorFunc: func(c *gin.Context) {
			c.String(403, "CSRF token mismatch")
			c.Abort()
		},
	}))

	// 限流
	router.Use(RateLimiter())

	middleware := router.Group("/middleware")

	{
		middleware.GET("/withoutMiddleware", WithoutMiddleware)
	}

	{
		middleware.GET("/usingMiddleware", UsingMiddleware)
	}

	// 下面这个是 分组中间件
	// middleware.Use(Logger())
	{
		// 这个则是 路由级中间件
		middleware.GET("/customMiddleware", Logger(), CustomMiddleware)
	}

	middleware.Use(ErrorHandler())
	{
		middleware.GET("/errorHandlingMiddleware", ErrorHandlingMiddleware)
	}

	// 使用 gin.BasicAuth() 中间件进行分组
	// gin.Accounts 是一个用于表示 map[string]string 的简写形式
	authorized := router.Group("/basicauth", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	{
		authorized.GET("/usingBasicauthMiddleware", UsingBasicauthMiddleware)
	}

	{
		middleware.GET("/goroutinesInsideAMiddleware", GoroutinesInsideAMiddleware)
	}

	{
		expectedHost := "localhost:8080"

		middleware.GET("/securityHeaders", func(c *gin.Context) {
			if c.Request.Host != expectedHost {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
				return
			}
			c.Header("X-Frame-Options", "DENY")
			c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
			c.Header("X-XSS-Protection", "1; mode=block")
			c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
			c.Header("Referrer-Policy", "strict-origin")
			c.Header("X-Content-Type-Options", "nosniff")
			c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
			c.Next()
		}, SecurityHeaders)
	}

	{
		middleware.GET("/securityGuide", SecurityGuide)
	}

	router.Run(":8080")
}
