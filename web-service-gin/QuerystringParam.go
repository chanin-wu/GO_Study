package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	c.Query("key") 返回查询参数的值，如果键不存在则返回空字符串。
	c.DefaultQuery("key", "default") 返回值，如果键不存在则返回指定的默认值。
*/

func QuerystringParam() {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":8080")
}
