package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("hello world")

	router := gin.Default()

	// 1. 搜索参数 ?key=value&key2=value2
	// 2. 路径参数 /path/:param

	// 搜索参数 ?key=value&key2=value2
	router.GET("/", func(c *gin.Context) {
		// c.Query("key") 获取搜索参数
		/* name := c.Query("name")
		age := c.Query("age") */

		// c.DefaultQuery("key","default value") 获取搜索参数，默认值
		name := c.DefaultQuery("name", "default name")
		age := c.DefaultQuery("age", "default age")
		c.String(http.StatusOK, "hello:%v world:%v", name, age)
	})

	// 路径参数 /path/:param
	router.GET("/detail/:id", func(c *gin.Context) {
		// c.Param("key") 获取路径参数
		id := c.Param("id")
		c.String(http.StatusOK, "hello:%v", id)

	})

	// POST 表单请求
	router.POST("/create", func(c *gin.Context) {
		/* name := c.PostForm("name")
		age := c.PostForm("age") */

		name := c.DefaultPostForm("name", "default name")
		age := c.DefaultPostForm("age", "default age")

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"name": name,
				"age":  age,
			},
			"success": true,
		})
	})

	// POST JSON请求
	router.POST("/create/json", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    user,
			"success": true,
		})
	})

	router.Run(":8080")
}
