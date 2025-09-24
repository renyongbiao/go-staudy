package main

import (
	"fmt"
	"time"
)

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
type Task func()

func executeTasks(tasks []Task) {

	for _, tast := range tasks {
		go func(t Task) {
			start := time.Now()
			t()
			elapsed := time.Since(start)
			fmt.Printf("任务执行时间: %s\n", elapsed)
		}(tast)
	}

}

func main() {
	// 任务调度

	tast1 := func() {
		fmt.Println("任务1执行了")
		time.Sleep(2 * time.Second)
	}
	tast2 := func() {
		fmt.Println("任务2执行了")
		time.Sleep(2 * time.Second)
	}
	executeTasks([]Task{tast1, tast2})

	time.Sleep(5 * time.Second) // 等待5秒，不然主协程终止导致子协程无法执行
}
