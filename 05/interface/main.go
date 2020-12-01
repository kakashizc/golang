package main

import "fmt"

type speaker interface {
	speak() //只要实现了speak()方法的变量,都是 speaker 类型
}
type cat struct {
}

type dog struct {
}

func (c cat) speak() {
	fmt.Println("苗苗")
}
func (d dog) speak() {
	fmt.Println("旺旺")
}

func out(x speaker) {
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	out(d1)
	out(c1)
}
