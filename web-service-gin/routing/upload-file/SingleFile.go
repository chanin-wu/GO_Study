package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SingleFile(c *gin.Context) {
	fmt.Println("SingleFile")
	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(file.Filename)

	// Upload the file to specific dst.
	dst := filepath.Join("./files/", filepath.Base(file.Filename))
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
