package main

import "fmt"

func main1() {

	// 数组
	a := [6]string{"1", "2", "3", "4", "5", "6"}

	fmt.Println(a)

	// 集合
	b := []string{"1", "2", "3", "4", "5", "6", "7"}

	fmt.Println(b[:4], b[4:])

	// b = append(b, "8", "9", "10") // 追加

	b = append(b, []string{"8", "9"}...) // 追加

	b = append([]string{"0"}, b...) // 前置

	// 第二个参数是可变参数，可以接收零个或多个 Type 类型的值。
	// 当使用 slice... 语法（即切片展开操作）时，它会将该切片的所有元素作为独立的参数依次传递给 append 的可变参数部分。该展开操作必须放在可变参数列表的末尾，因为展开后会产生不定数量的参数，编译器无法再确定后续的参数应该以何种方式解析。
	// b = append(b[:4],[]string{"8", "9"}..., "10") 写法错误
	b = append(b[0:4], append([]string{"8", "9"}, b[4:]...)...) // 中间

	fmt.Println(b)
}
