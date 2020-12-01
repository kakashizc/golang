package main

import "fmt"

//闭包
func main() {
	a := addr(100)
	b := a(200)
	fmt.Println(b)

	f2(f1)
}

//闭包是一个函数,这个函数 包含了他外部作用域的一个变量
//底层原理 1, 函数可以作为一个返回值
//2,函数内部查找变量的顺序,现在自己内部查找,找不到往外层找
func addr(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func f1(int, int) {

}

func f2(f func(int, int)) func() {
	return func() {

	}
}
