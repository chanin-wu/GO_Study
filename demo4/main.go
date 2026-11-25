package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	中间件
		跨域
		限校验
		IP限流
		登录鉴权
		黑名单
		日期请求记录
		响应记录
		耗时统计
*/

/* func Cors(c *gin.Context) {
	fmt.Println("Cors Before 业务之前处理")
	c.Next()  // 当我们不去调用Next() 方法时，后续的业务处理便不会执行
	c.Abort() // 中断后续所有链条处理
	fmt.Println("Cors After 业务之后处理")
} */

func Cors2(c *gin.Context) {
	fmt.Println("Cors Before2")

	// 高级请求 POST+json DELETE PUT PATCH 等
	// 普通请求 GET POST+表单方式 两种表单form+url
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")                                       // 允许所有域名访问
	c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE, PATCH") // 允许所有方法访问
	c.Writer.Header().Set("Access-Control-Allow-Headers", "content-type")                           // 允许所有头访问访问
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "*")                                  // 允许所有凭证访问

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}

	c.Next()
	fmt.Println("Cors After2")
}

func main() {
	// Default() 创建一个路由 其中已经添加了跨域中间件
	// router := gin.Default()

	// New() 创建一个纯净的路由 再调用 Use() 添加中间件
	router := gin.New()

	// 代码自上而下执行 中间件 要放在业务处理之前 洋葱模型
	// router.Use(Cors)
	router.Use(Cors2)

	router.GET("/ping", func(c *gin.Context) {
		fmt.Println("Ping Run")
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.Run(":8080")
}
