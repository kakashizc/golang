package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex
var wrlock sync.RWMutex

//读写锁  读的时候都可以读, 只有写的时候锁住
func read() {
	//lock.Lock()
	wrlock.RLock() //换成Rlock() 执行效率会高很多,实测 lock 跑完程序需要1.2s , Rlock用了0.01s
	defer wg.Done()
	fmt.Println(x)
	time.Sleep(time.Millisecond) //假设读操作需要1毫秒
	wrlock.RUnlock()
	//lock.Unlock()
}

func write() {
	//lock.Lock()
	wrlock.RLock()
	defer wg.Done()
	x = x + 1
	fmt.Println("x加1了...=", x)
	time.Sleep(time.Millisecond * 5) //假设读操作需要5毫秒
	wrlock.RUnlock()
	//lock.Unlock()
}

//模拟互斥锁 和 读写锁的性能差异
func main() {
	start := time.Now()
	//执行10次写操作
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}
	//执行10000次 读操作, 因为实际应用中 读的场景大大于写的场景
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
