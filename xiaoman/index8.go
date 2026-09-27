package main

import (
	"fmt"
	"time"
)

/*
	chan 是队列吗？
	一、语义上是，实现上分情况：
		1.带缓冲的 channel（make(chan T, n)）：就是一个 FIFO 队列。底层是运行时里的一个环形缓冲区（ring buffer），发送入队尾、接收出队头，顺序严格先进先出。
		2.无缓冲的 channel（make(chan T)）：没有任何存储，发送方会阻塞到有接收方就绪，两者直接"交接"（同步点），更像一个会合点而非队列。
		3.当缓冲满/空时，阻塞的 goroutine 会挂在 sendq/recvq 等待队列里，也是 FIFO 唤醒的。

	二、能不能只取第几个元素？
		不能。 channel 只支持从队头取，取走即删除：
		1.没有随机访问，不能 ch[3]
		2.没有 peek（偷看不消费）
		3.想"跳过前几个"只能逐个读出来丢弃，它们照样被消费掉了

	这是刻意设计的——channel 是协程间通信和同步的原语，不是通用容器。
*/

// 子协程 无缓冲通道
func no_buffer_worker(ch chan int) {
	ch <- 100
	ch <- 200
	ch <- 300

	// 通道中请添加 close
	close(ch)
}

// 子协程 无缓冲通道
func no_buffer_worker1(ch chan int) {
	// 运行机制 双向奔赴 子协程发送 主协程必须接受
	fmt.Println("无缓冲通道start")
	// 子协程发送数据 如果主协程 没有接受 就会阻塞后续代码执行
	ch <- 400
	fmt.Println("无缓冲通道end")
}

// 主协程
func main8() {
	// 1. 无缓冲通道 不定义容量
	// 2. 有缓冲通道 定义容量
	ch := make(chan int) // 无缓冲通道
	go no_buffer_worker(ch)

	/* val1 := <-ch
	val2 := <-ch
	val3 := <-ch

	fmt.Println(val1, val2, val3) */

	// for循环类似于一个水闸 他会等待这个通道一直传递数据 不会自动关闭水闸 控制会报错 deadlock! 死锁
	for val := range ch {
		fmt.Println(val)
	}

	// ------------------------------------------------

	ch1 := make(chan int) // 无缓冲通道
	go no_buffer_worker1(ch1)
	// time.Sleep(1 * time.Second)
	val4 := <-ch1
	fmt.Println(val4)

	// ------------------------------------------------

	// 原理：make() 第二个参数的数字 代表 这个子协程能处理几个通道 为2 只能处理2个 超出的会产生死锁
	// 解决方案：在下方消费掉即可
	// [100,200] 300
	ch2 := make(chan int, 2)
	go with_buffer_worker(ch2)
	val5 := <-ch2
	fmt.Println(val5)
	// time.Sleep(1 * time.Second)

	// ------------------------------------------------

	work1 := make(chan string)
	go worker1(work1)

	work2 := make(chan string)
	go worker2(work2)

	select {
	case msg1 := <-work1:
		fmt.Println(msg1)
	case msg2 := <-work2:
		fmt.Println(msg2)
	case <-time.After(20 * time.Millisecond):
		fmt.Println("超时展示数据啦")
	default:
		fmt.Println("非阻塞测试")
	}

}

// 有缓冲通道
func with_buffer_worker(ch chan int) {
	fmt.Println("有缓冲通道1111")
	ch <- 100
	ch <- 200
	ch <- 300
	fmt.Println("有缓冲通道2222")
}

/*
	执行机制 如果有多个通道 优先返回延迟最低的那一个
	select 用于监听多个channel
	go 的 select 用于处理并发竞争和超时
	1、多channel同时监听
	2、请求超时控制
	3、非阻塞读写channel
*/

// 如果下方两个 时间 相同 time.Sleep(30 * time.Millisecond)
// 那么则会随即返回
func worker1(ch chan string) {
	time.Sleep(50 * time.Millisecond)
	ch <- "50毫秒 v1"
}

func worker2(ch chan string) {
	time.Sleep(30 * time.Millisecond)
	ch <- "30毫秒 v2"
}
