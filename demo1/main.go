package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("hello world")

	// http.NewServeMux() 创建多路 路由复用器 路由实例独立 不会被污染
	// http.HandleFunc("/hello") 创建公开的路由 可能会被污染

	mux := http.NewServeMux()

	// 定义独立路由
	mux.HandleFunc("GET /user/list", func(w http.ResponseWriter, r *http.Request) {
		// w Response => 返回资源
		// r Request => 请求资源
		fmt.Println("user list")

		// 切片语法 : []byte{}
		// 语法糖 : []byte(类型) () => []byte{}
		// 支持三种转换的格式 字符串 | 字节切片 | nil

		// w.Write([]byte("user list"))

		userList := []User{
			{
				ID:   1,
				Name: "张三",
			},
			{
				ID:   2,
				Name: "李四",
			},
			{
				ID:   3,
				Name: "王五",
			},
		}
		targetId := r.URL.Query().Get("id")
		fmt.Println(targetId)

		if targetId != "" {
			for _, v := range userList {
				// strconv.Itoa() 整数转字符串
				if targetId == strconv.Itoa(v.ID) {
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					json.NewEncoder(w).Encode(v)
					return
				}
			}
		}

		// 默认返回MIME类型为text/plain
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(userList) // 将userList转换为JSON字符串并写入w

	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe() // 启动服务器
}
