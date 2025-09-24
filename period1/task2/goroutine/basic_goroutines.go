package main

import (
	"fmt"
	"time"
)

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 打印奇数
func printOdd() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数:", i)
	}
}

// 打印偶数
func printEven() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数:", i)
	}
}

func main() {
	//协程
	go printOdd()
	go printEven()
	time.Sleep(1 * time.Second) // 等待1秒，不然主协程终止导致子协程无法执行
}
