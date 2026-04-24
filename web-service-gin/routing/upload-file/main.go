package main

import "github.com/gin-gonic/gin"

/*
Gin 为 multipart 表单解析设置了默认 32 MiB 的内存限制，通过 router.MaxMultipartMemory 设置。
在此限制内的文件会缓存在内存中；超出的部分会写入磁盘上的临时文件。
*/

func main() {
	router := gin.Default()

	router.POST("single-file", SingleFile)
	router.POST("multiple-file", MultipleFile)
	router.POST("limit-bytes", LimitBytes)

	router.Run(":8080")
}
