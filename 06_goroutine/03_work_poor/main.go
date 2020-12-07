package main

import (
	"fmt"
	"time"
)

//这里 jobs 单向通道只能取值
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second * 3)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启3个goroutine 去执行下面的5个任务, worker1 woker2  woker3  各取一个任务,谁先执行完, 再去请求剩下两个任务中的一个
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)

	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j //这里josb 并不是单向通道, 所以可以往它里面存值
	}
	close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
