package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test(a int) {
	fmt.Println(a)
}

func main() {
	a := 12
	go test(a)
	time.Sleep(time.Second)
}
