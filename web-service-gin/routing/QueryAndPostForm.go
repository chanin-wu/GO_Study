package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	c.Query("key") / c.DefaultQuery("key", "default") —— 从 URL 查询字符串读取。
	c.PostForm("key") / c.DefaultPostForm("key", "default") —— 从 application/x-www-form-urlencoded 或 multipart/form-data 请求体读取。
*/

func QueryAndPostForm() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.DefaultPostForm("message", "666")

		fmt.Print("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
		c.String(http.StatusOK, "id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	router.Run(":8080")
}
