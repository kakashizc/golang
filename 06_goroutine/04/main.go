package main

import "fmt"

//range取值
func main() {

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go func() {

		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			v, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- v
		}
		close(ch2)
	}()

	for value := range ch2 {
		fmt.Println(value)
	}

}
