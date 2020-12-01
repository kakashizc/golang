package main

import "fmt"

type animal interface {
	eat(string)
	run(speed int8)
}

type dog struct {
	feet int8
}

func (d dog) eat(t string) {
	fmt.Printf("%s,真棒...", t)
}
func (d dog) run(s int8) {
	fmt.Printf("狗跑的真快都%dkm了", s)
}

type chicken struct {
	feet int8
}

func (c chicken) eat(t string) {
	fmt.Printf("%s鸡吃东西...", t)
}

func main() {
	var a animal //定义一个接口类型的变量

	d1 := dog{
		feet: 4,
	}
	a = d1
	fmt.Println(a)
	a.run(11)
	// a = c1
	// fmt.Println(a)
	// a.eat("大黄送")
}
