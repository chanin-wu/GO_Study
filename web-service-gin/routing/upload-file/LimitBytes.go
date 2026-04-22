package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MaxUploadSize = 1 << 20 // 1 MB
)

func LimitBytes(c *gin.Context) {
	fmt.Println("LimitBytes")

	// Wrap the body reader so only MaxUploadSize bytes are allowed
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxUploadSize)

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
	defer file.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successful",
	})
}
