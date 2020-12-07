package main

import (
	"fmt"
	"sync"
)

//channel通道联系
//1,启动一个goroutine,生成100个数 发送到通道ch1
//2,启动一个goroutine,从ch1中取值,起算起平方  放到ch2中
//3, 在main中, 把ch2的值打印出来
var wg sync.WaitGroup

//1,生成100个数 发到通道ch1
func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1) //关闭通道 不能存 但是能取值
}

func f2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok { //说明没有值,break
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(2)
	go f1(a)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}
