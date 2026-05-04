package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	rendering := router.Group("/rendering")

	{
		rendering.GET("/someJSON", SomeJSON)
		rendering.GET("/someXML", SomeXML)
		rendering.GET("/someYAML", SomeYAML)
		rendering.GET("/someProtoBuf", SomeProtoBuf)
	}

	{
		// You can also use your own secure json prefix
		// router.SecureJsonPrefix(")]}',\n")
		rendering.GET("/secureJson", SecureJson)
	}

	{
		rendering.GET("/json", Json)
		rendering.GET("/purejson", PureJson)
	}

	router.Run(":8080")
}
