package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ch1 chan bool

//context初识
//为什么需要context
/*
比如我们开启了一个副goroutine 想要在main主goroutine中关闭 副goroutine 如何实现?
1, 定义一个全局变量 在mian()中控制时间 改变变量的值 ,for中监听到变量的值改变以后 break
2, 定义一个channel  for中监听
3, 用go新加入的 context 来做, (其实原理就是channel值的改变)
*/
func f(c chan bool) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("副goroutine打印结果")
		time.Sleep(time.Millisecond * 500) //500毫秒 打印一次 0.5秒
		select {
		case <-c:
			//break //此次跳出的是select 循环,而不是for循环,所以要用指定跳出循环
			break LOOP
		default:
		}
	}
}

func main() {
	ch1 := make(chan bool, 1)
	wg.Add(1)
	go f(ch1)
	time.Sleep(time.Second * 5) //停5秒
	ch1 <- true
	wg.Wait()
}
