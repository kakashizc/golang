package main

import (
	"fmt"
	"sync"
)

//互斥锁  如下方式会导致两个goroutine 同时去操作x变量
//比如: 第一个goroutine获取的x是50的时候 ,第二个也同时获取了50 他俩都+1 x加了两次应该是52,但是此时是51了,这样就少加了一个数,造成冲突
//所以需要用互斥锁 类似于mysql中的读锁 , 当我读的时候lock 不让别人动.  不过这样性能低下, 具体解决请查阅下个单元
var x = 0
var wg sync.WaitGroup
var lock sync.Mutex //引入锁
func add() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		lock.Lock() //加锁
		x = x + 1
		lock.Unlock() //释放锁
	}
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
