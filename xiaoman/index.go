package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	互斥锁 VS 读写锁
	1、读和读互斥		读和读共享
	2、写和写互斥		写和写互斥
	3、读和写互斥		读和写互斥
*/

func worker(nums *[]int, num int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	// 临界区： Lock 和 Unlock 之间的代码称作为临界区
	*nums = append(*nums, num)
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex     // 互斥锁 同一时间内只能有一个协程去操作临界区的共享数据
	var rwmu sync.RWMutex // 读写锁 适用于读多写少
	nums := []int{1, 2, 3}
	wg.Add(6)
	// 互斥锁 演示
	go worker(&nums, 4, &wg, &mu) // 变成4个
	go worker(&nums, 5, &wg, &mu) // 变成4个
	go worker(&nums, 6, &wg, &mu) // 变成4个 覆盖操作 有几率变成6个

	// 读写锁演示
	go read(&nums, &wg, &rwmu)
	go read(&nums, &wg, &rwmu)
	go write(&nums, 7, &wg, &rwmu)
	wg.Wait()
	fmt.Println(nums)
}

func read(nums *[]int, wg *sync.WaitGroup, rwmu *sync.RWMutex) {
	defer wg.Done()
	rwmu.Lock()
	time.Sleep(1 * time.Millisecond) // 1毫秒延迟
	fmt.Println("read", nums)
	rwmu.Unlock()
}

func write(nums *[]int, num int, wg *sync.WaitGroup, rwmu *sync.RWMutex) {
	time.Sleep(50 * time.Millisecond) // 50毫秒延迟
	defer wg.Done()
	rwmu.Lock()
	*nums = append(*nums, num)
	rwmu.Unlock()
}
