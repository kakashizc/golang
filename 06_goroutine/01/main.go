package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	runtime.GOMAXPROCS(4) //设置几个cpu 核心去执行任务 默认是最多
	wg.Add(2)             //等级一下,需要用到几个goroutine
	go a()
	go b()
	wg.Wait() //计数器归零,程序结束
}

func a() {
	defer wg.Done() //用完了 要 -1
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done() //用完了 -1
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}
