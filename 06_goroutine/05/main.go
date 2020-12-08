package main

import "fmt"

func worker(c chan int) {
	//从channel中去读数据
	num := <-c
	fmt.Println("foo recv channel ", num)
}

func main() {
	//创建一个channel
	c := make(chan int)

	go worker(c) //worker启动后 一直在等待通道 c 被赋值,  被赋值后num就有值了,就可以打印输出了
	//main协程 向一个channel中写数据
	c <- 2
	fmt.Println("send 1 -> channel over")
}
