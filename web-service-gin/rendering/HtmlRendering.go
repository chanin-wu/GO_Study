package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
使用 LoadHTMLGlob() 或 LoadHTMLFiles() 来选择要加载的 HTML 文件。
*/

func HtmlRendering(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}

func PostsIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
		"title": "Posts",
	})
}
