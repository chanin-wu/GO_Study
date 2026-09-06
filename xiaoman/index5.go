package main

import (
	"fmt"
)

func MultiplicationTable(n int) (bool, error) {
	if n != 9 {
		panic("参数错误") // 程序异常 会阻断后续的代码执行
		// return false, errors.New("参数错误")
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}

		fmt.Println()
	}

	return true, nil
}

func main5() {
	// 这里可以捕获 程序异常 但必须要添加defer
	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	status, err := MultiplicationTable(9)
	if !status {
		fmt.Println(status, err)
	}

	// 业务异常 不影响后续代码执行
	fmt.Println(1)       // 优先执行
	defer fmt.Println(2) // 把后面的代码推迟到函数末尾去执行
	defer fmt.Println(3) // 把后面的代码推迟到函数末尾去执行
	defer fmt.Println(4) // 把后面的代码推迟到函数末尾去执行
	fmt.Println(5)       // 优先执行
	// defer (后进先出)
	// 普通代码先执行 -> return赋值给返回值 -> defer -> return返回
	// 1 5 4 3 2
}

/* Go 里这三个打印函数，新人最容易混，记住一个口诀就行：
Print 是裸奔，Printf 是填空，Println 是换行。

三个的区别
1. fmt.Print()
原样输出，不换行、不加空格。

fmt.Print("hello")
fmt.Print("world")
// 输出：helloworld

2. fmt.Println()
输出 + 自动换行，多个参数之间用空格隔开。

fmt.Println("hello", "world")
// 输出：hello world\n

3. fmt.Printf()
格式化输出，不自动换行。第一个参数是格式字符串，用占位符填值。

fmt.Printf("name=%s, age=%d\n", "Mavis", 18)
// 输出：name=Mavis, age=18 */

/* 常用占位符（Printf 专属）
占位符	含义	示例
%v	万能占位，啥都能打	Printf("%v", user)
%+v	结构体带字段名	Printf("%+v", user)
%T	打印类型	Printf("%T", x)
%s	字符串	Printf("%s", "hi")
%d	整数	Printf("%d", 18)
%f	浮点数	Printf("%.2f", 3.1415) → 3.14
%t	布尔	Printf("%t", true)
%p	指针	Printf("%p", &x)
%q	带引号的字符串	Printf("%q", "hi") → "hi" */

/*
实战小建议
调试大结构体：直接 fmt.Printf("%+v\n", obj)，字段名和值一起出来，比一个个 Print 字段快多了
错误位置定位：fmt.Printf("file=%s line=%d\n", file, line)
性能敏感场景：fmt.Fprintf(os.Stdout, ...) 直接写文件描述符，比包一层 Printf 快
*/

/*
	panic + defer + recover 是 Go 的"三件套"
	先说 panic 的传播机制
		panic 抛出后，会沿着调用栈一路往上抛，像滚雪球一样：
		main()
			└─ MultiplicationTable()  ← panic 在这里产生
					└─ for loop
								└─ panic("参数错误")

	panic 会做两件事：
	1.立刻停止当前函数的执行
	2.一边往上抛，一边执行每一层 defer 注册的函数（这就是 defer 的关键作用）

	为什么 recover 必须在 defer 里？
	这是 Go 的语言级硬性规定——recover() 只有在 defer 函数里调用才有效。在普通函数里调用，永远返回 nil。

	串起来：为什么必须 defer？
	把整个执行流程走一遍你的代码（传入 n=10）：
	第 1 步：main() 开始执行
        注册 defer 函数（注意：此时 defer 函数还没执行，只是"挂了个钩子"）
	第 2 步：调用 MultiplicationTable(10)
					走到 if n != 9 时触发 panic("参数错误")
	第 3 步：panic 触发后：
					- 立刻停止 MultiplicationTable 后续代码
					- 开始沿调用栈往上"展开"
					- 展开过程中，会执行每一层注册过的 defer 函数
	第 4 步：MultiplicationTable 自身没有 defer，展开到 main
	第 5 步：main 里的 defer 函数执行
					- recover() 捕获到 panic("参数错误")
					- 程序不再崩溃，从这里继续往下走
	第 6 步：defer 之后，main 里剩下的代码继续执行
					（你的代码里没有剩余代码，所以程序正常结束）

	一句话总结
	defer 是"抢救现场的窗口"，panic 是"正在发生的灾难"，recover 是"在窗口里伸出的手"。

	顺手补一个对比
	机制	作用	类比
	panic	主动抛出异常，停止当前函数	火灾报警器响了
	defer	注册一个"无论如何最后都会执行"的函数	火灾应急预案
	recover	在 defer 里调用，捕获 panic，控制权返回	应急预案里的灭火器

	实战建议
	recover 只在两个地方用：

	1.最外层（如 main、gin 的 recovery 中间件）：把 panic 兜住，避免进程整个挂掉
	2.包/库的边界：防止内部 panic 把外部搞崩
*/
