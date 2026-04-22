package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParamInPath() {
	router := gin.Default()

	// :name —— 匹配单个路径段。例如，/user/:name 匹配 /user/john，但不匹配 /user/ 或 /user。
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// *action —— 匹配前缀之后的所有内容，包括斜杠。例如，/user/:name/*action 匹配 /user/john/send 和 /user/john/。捕获的值包含前导 /。
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
