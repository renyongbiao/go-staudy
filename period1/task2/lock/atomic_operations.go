package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter1 int64 = 0 // 原子操作要求 int64 类型
	var wg3 sync.WaitGroup

	wg3.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg3.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter1, 1) // 原子递增
			}
		}()
	}
	wg3.Wait()
	fmt.Println("final counter:", counter1) // 预期 10000
}
