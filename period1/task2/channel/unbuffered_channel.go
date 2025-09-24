package main

import (
	"fmt"
	"sync"
)

func main() {
	// 协程 无缓冲
	parmChannel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			parmChannel <- i
		}
		close(parmChannel)
	}()
	go func() {
		defer wg.Done()
		// 接收参数
		for v := range parmChannel {
			fmt.Println("接收的参数为：", v)
		}
	}()
	wg.Wait()
}
