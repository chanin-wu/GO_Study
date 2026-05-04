package main

import (
	"github.com/gin-gonic/gin"
)

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

	{
		rendering.GET("/servingStaticFiles", ServingStaticFiles)
		// router.Static("/static", "./static")
		// router.StaticFS("/static", http.Dir("static"))
		// router.StaticFile("/static/text.txt", "./static/text.txt")
	}

	{
		rendering.GET("/local/file", LocalFile)
		rendering.GET("/fs/file", FsFile)
		rendering.GET("/download", Download)
	}

	{
		rendering.GET("/servingDataFromReader", ServingDataFromReader)
	}

	{
		router.LoadHTMLGlob("templates/*")
		rendering.GET("/htmlRendering", HtmlRendering)

		// router.LoadHTMLGlob("templates/**/*")
		// rendering.GET("/posts/index", PostsIndex)
	}

	router.Run(":8080")
}
