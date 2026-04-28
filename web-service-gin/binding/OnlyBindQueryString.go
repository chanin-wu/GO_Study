package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
ShouldBindQuery 仅将 URL 查询字符串参数绑定到结构体，完全忽略请求体。当你需要确保 POST 请求体数据不会意外覆盖查询参数时——例如在同时接受查询过滤器和 JSON 请求体的端点中——这将非常有用
*/

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func OnlyBindQueryString(c *gin.Context) {
	var person Person
	if err := c.ShouldBindQuery(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":    person.Name,
		"address": person.Address,
	})
}
