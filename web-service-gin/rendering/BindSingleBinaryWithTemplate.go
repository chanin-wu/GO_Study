package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	自 Go 1.16 起，标准库支持使用 //go:embed 指令将文件直接嵌入二进制文件中。无需第三方依赖
*/

func BindSingleBinaryWithTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}
