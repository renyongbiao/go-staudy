package main

import (
	"fmt"
	"sync"
)

func main() {
	// 协程  有缓冲

	ch := make(chan int, 10) // ① 带缓冲的通道
	var wg1 sync.WaitGroup
	wg1.Add(2)

	// 生产者
	go func() {
		defer wg1.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch) // 发完关闭
	}()

	// 消费者
	go func() {
		defer wg1.Done()
		for v := range ch { // 通道关闭后自动退出
			fmt.Println(v)
		}
	}()

	wg1.Wait() // 等双方结束
}
