package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AsciiJSON() {
	router := gin.Default()
	router.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})

	router.Run("localhost:8080")
}
