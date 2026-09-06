package main

import "fmt"

// 闭包
func closure() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main3() {
	next := closure()
	fmt.Println(next())
	fmt.Println(next())
}
