package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func isAllowedExtension(allowExtensions []string, ext string) bool {
	for _, e := range allowExtensions {
		if e == ext {
			return true
		}
	}
	return false
}

func SingleFile(c *gin.Context) {
	fmt.Println("SingleFile")

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// 10 表示 kb
	// 20 表示 mb
	// 30 表示 gb
	// router := gin.Default()
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// 限制上传文件大小为 8 MB
	const MaxUploadSize = 8 << 20 // 8 MiB

	// 允许上传的文件扩展名
	var allowExtensions = []string{".png", ".jpg", ".jpeg"}

	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(file.Filename)

	// 检查文件大小是否超过限制
	if file.Size > MaxUploadSize {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": fmt.Sprintf("file too large (max: %d bytes)", MaxUploadSize)})
		return
	}

	// 检查文件扩展名是否在允许列表中
	if !isAllowedExtension(allowExtensions, filepath.Ext(file.Filename)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file extension not allowed"})
		return
	}

	// Upload the file to specific dst.
	dst := filepath.Join("./files/", filepath.Base(file.Filename))
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
