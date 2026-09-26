package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	Name string        // 订单名称
	Time time.Duration // 模拟延迟网络请求的属性
}

func notify(order Order, wg *sync.WaitGroup) {
	defer wg.Done() // 表示子协程执行完毕， Done() 必须要放在函数执行之前，执行完毕一个 减去一个 count - 1

	time.Sleep(order.Time)
	fmt.Println(order.Name)
}

// 主协程
/*
	主协程代码执行完毕后就会立即退出
	主协程退出后 子协程 无论有没有执行完毕 都会退出
*/
func main7() {
	var wg sync.WaitGroup
	orders := []Order{
		{Name: "1", Time: 1000 * time.Millisecond},
		{Name: "2", Time: 2000 * time.Millisecond},
		{Name: "3", Time: 3000 * time.Millisecond},
		{Name: "4", Time: 4000 * time.Millisecond},
	}

	// 告诉 go语言 有几个子协程
	wg.Add(len(orders))
	for _, order := range orders {
		go notify(order, &wg)
	}
	wg.Wait() // 让主协程等待， 原理：for死循环 if count === 0 才会退出

	fmt.Println("主协程执行完毕")
}
