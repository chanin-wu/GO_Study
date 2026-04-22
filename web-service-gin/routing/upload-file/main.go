package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("single-file", SingleFile)
	router.POST("multiple-file", MultipleFile)
	router.POST("limit-bytes", LimitBytes)

	router.Run(":8080")
}
