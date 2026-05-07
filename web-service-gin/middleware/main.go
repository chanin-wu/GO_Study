package main

import (
	"net/http"
	"time"

	"github.com/danielkov/gin-helmet/ginhelmet"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/utrack/gin-csrf"
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
		AllowOrigins:     []string{"https://example.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// CSRF
	store := cookie.NewStore([]byte("session-secret"))
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
