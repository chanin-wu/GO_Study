package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	c.QueryMap("key") —— 从 URL 查询字符串中解析 key[subkey]=value 形式的键值对。
	c.PostFormMap("key") —— 从请求体中解析 key[subkey]=value 形式的键值对。
*/

/* curl -X POST "http://localhost:8080/post?ids[a]=1234&ids[b]=hello" \
-d "names[first]=thinkerou&names[second]=tianou" */

func MapAsQuerystringOrPostform() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Print("ids: %v; names : %v\n", ids, names)
		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	router.Run(":8080")
}
