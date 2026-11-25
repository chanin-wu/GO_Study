package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
定义限制 —— 常量 MaxUploadSize（1 MB）设置上传的硬性上限。
强制限制 —— http.MaxBytesReader 包装 c.Request.Body。如果客户端发送的字节数超过允许值，读取器将停止并返回错误。
解析并检查 —— c.Request.ParseMultipartForm 触发读取。代码检查 *http.MaxBytesError 以返回带有明确消息的 413 状态码。
*/

const (
	MaxUploadSize = 1 << 10 // 1 KB
)

func LimitBytes(c *gin.Context) {
	fmt.Println("LimitBytes")

	// 将该读取器进行封装，使其仅允许读取最多“MaxUploadSize”字节的数据。
	// Wrap the body reader so only MaxUploadSize bytes are allowed
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxUploadSize)

	// 解析多部分表单
	// Parse multipart form
	if err := c.Request.ParseMultipartForm(MaxUploadSize); err != nil {
		if _, ok := err.(*http.MaxBytesError); ok {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"error": fmt.Sprintf("file too large (max: %d bytes)", MaxUploadSize),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file form required"})
		return
	}
	// defer 关键字用于注册一个延迟执行的函数调用
	defer file.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successful",
	})
}
