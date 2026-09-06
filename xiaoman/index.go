package main

import "fmt"

// 值传递
func change(x *int) {
	// 修改的是一个副本 不是原始数据
	// 使用的时候需增加*
	*x = 100
}

// 下面是引用类型
// 集合 切片传递
func change1(a []int) {
	// go 语言在设计切片的时候 它里面的数据就是一个指针
	a[0] = 4
	// 切片s本身还是一个副本 添加的时候 并不会生效
	a = append(a, 5)
}

// 数组传递 数组本身是一个(值传递)
func change2(arr [3]int) {
	arr[0] = 4
}

// map传递
func change3(m map[string]int) {
	m["语文"] = 4
	delete(m, "英语")
}

type Person struct {
	Name string
	Age  int
	Sex  bool
}

// 结构体传递 结构体本身是一个(值传递)
func change4(p *Person) {
	p.Age = 200.     // 简写
	(*p).Sex = false // 全称
}

func main() {
	x := 10
	// 指针
	// &x 取地址
	// *x 解引用
	change(&x)
	fmt.Println("值传递", x, &x)

	a := []int{1, 2, 3}
	change1(a)
	fmt.Println("集合传递", a)

	arr := [3]int{1, 2, 3}
	change2(arr)
	fmt.Println("数组传递", arr)

	m := map[string]int{
		"语文": 1,
		"数学": 2,
		"英语": 3,
	}
	change3(m)
	fmt.Println("map传递", m)

	p := Person{
		Name: "cigarette",
		Age:  18,
		Sex:  true,
	}
	change4(&p)
	fmt.Println("结构体传递", p)
}
