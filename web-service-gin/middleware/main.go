package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	middleware := router.Group("/middleware")

	{
		middleware.GET("/withoutMiddleware", WithoutMiddleware)
	}

	{
		middleware.GET("/usingMiddleware", UsingMiddleware)
	}

	middleware.Use(Logger())
	{
		middleware.GET("/customMiddleware", CustomMiddleware)
	}

	router.Run(":8080")
}
