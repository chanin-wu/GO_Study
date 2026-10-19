package main

import (
	"fmt"
	"os"
)

func basics() {
	// 创建文件
	file, _ := os.Create("test.txt")

	// 追加内容、写入内容
	file.WriteString("hello word 1\n")

	// 追加内容、写入内容
	file.Write([]byte("hello word 2"))

	file.Close()

	// 读取文件
	data, _ := os.ReadFile("test.txt")
	fmt.Println(string(data))

	// 获取文件信息
	fileInfo, _ := os.Stat("test.txt")
	fmt.Println(fileInfo)

	os.Remove("test.txt")
}

func advance() {
	// https://go-docs.pages.dev/posts/fs/
	// 打开文件
	file, err := os.OpenFile("advance.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// 追加内容、写入内容
	file.WriteString("hello word 1\n")
	file.Write([]byte("hello word 2\n"))

	/*
		0o777 windows无效 只针对 linux/macOS

		0o 八进制

		三个7 代表代表的含义
		7 文件的所有者
		7 用户组
		7 其他人

			7 = 4 2 1

			4 读取 r
			2 写入 w
			1 执行 x

			5 读取 r + 执行 x
			6 读取 r + 写入 w
	*/
}

func main10() {
	// basics()
	advance()
}
