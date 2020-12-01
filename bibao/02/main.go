package main

import "fmt"

//闭包的实际应用

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//要求 f1(f2) ,但是当前f1 的参数是匿名函数,而f2是有名函数, 这样就需要另外一个函数f3转接一下了

func f3(f func(int, int), x, y int) func() {
	//1,此函数的返回值是一个匿名函数, 就可以作为f1的参数传给f1了
	tmp := func() {
		//fmt.Println(x + y)
		//2,调用f3方法的时候传入f2函数, 在这调用参数1 f相当于调用f2函数了, 实现了我们的需求
		f(x, y) //执行第一个参数,也就是f2函数
	}
	return tmp
}

func main() {
	f1(f3(f2, 100, 200))
}
