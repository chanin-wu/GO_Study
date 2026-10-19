package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Human struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address *string `json:"address"` // 设置为指针类型，表示地址可以为空 nil
}

// 序列化 -> Go语言数据格式转换 -> JSON数据格式
// 反序列化 -> JSON数据格式 -> Go语言数据格式转换
func main() {
	person := Human{
		Name: "xiaoman",
		Age:  18,
	}
	// 转换为 JSON 字符串
	jsonStr, _ := json.Marshal(person)
	fmt.Println(string(jsonStr))

	// 从 JSON 字符串转换为结构体
	person2 := []Human{}
	data, _ := os.ReadFile("data.json")
	json.Unmarshal(data, &person2)
	fmt.Println(person2)
	fmt.Printf("%+v\n", person2[0])
	fmt.Printf("%s\n", *person2[0].Address)
}

/* ## 原因分析
1. `*string` 是指针类型 。 index.go 中`Address *string` 存的是"指向字符串的地址"，而不是字符串本身。
2. `fmt.Println` 打印结构体时，指针字段只打印地址值 。所以`person2` 里的`Address` 显示为`0x20caed42e0f0` 这种十六进制指针值。
3. 每次运行地址都不同 （你看到的是`0x2c196f4220f0` ，我这次是`0x20caed42e0f0` ）——这正说明它是内存地址，而非编码错乱的文本。如果是乱码，内容会是固定的怪异字符，而不是每次变化的`0x...` 。
4. 反序列化本身是成功的 ：`json.Unmarshal` 已把`"杭州"` 正确写入堆内存，`Address` 指针指向的就是"杭州"，只是打印时没解引用。
如果你只是想让`Println` 直接显示内容，把字段改回`string` 即可；保留`*string` 的价值在于能区分 JSON 中`"address": null` （字段存在但为空）和字段缺失（`nil` ）两种情况——这正是你注释里写的"可以为空 nil"的用途，只是打印方式需要相应调整。 */
