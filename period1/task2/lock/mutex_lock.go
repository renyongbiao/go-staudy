package main

import (
	"fmt"
	"sync"
)

func main() {
	//锁机制
	var (
		counter int        // 共享计数器
		mu      sync.Mutex // 保护 counter 的互斥锁
		wg2     sync.WaitGroup
	)

	wg2.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg2.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock() // 加锁
				counter++
				mu.Unlock() // 解锁
			}
		}()
	}
	wg2.Wait()
	fmt.Println("final counter:", counter) // 预期 10000
}
