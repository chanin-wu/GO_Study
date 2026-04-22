package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	c.PostForm("field") 返回值，如果字段不存在则返回空字符串。
	c.DefaultPostForm("field", "fallback") 返回值，如果字段不存在则返回指定的默认值。
*/

func MultipartUrlencodedForm() {
	router := gin.Default()

	router.POST("form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.Run(":8080")
}
