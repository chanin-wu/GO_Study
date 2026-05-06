package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()

	// 下面两个是 全局中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

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

	router.Run(":8080")
}
