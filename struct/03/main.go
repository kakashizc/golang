package main

import "fmt"

type person struct {
	name string
	age  int
}

//构造函数
func newPerson(name string, age int) person {
	return person{
		name,
		age,
	}
}

//方法是作用于特定类型的函数
func (p person) add() {
	p.age++
	fmt.Println("方法结构体的age:", p.age) //13
}

//指针接收者(注意:只要有一个方法使用了指针接收者,其他的方法也应该是指针接收者,这里为了测试 add 方法就不改了)
func (p *person) readd() {
	p.age++
}

func main() {
	a := newPerson("黄大", 12)
	a.add()
	fmt.Println("原来的结构体的age:", a.age) // 12 以前的结构体的值不变,因为dog()方法中的person结构体是复制本,他有一个独立的内存地址,他的改变不会改变原结构体
	a.readd()
	fmt.Println("原来的结构体的age:", a.age) // 13 readd()方法用的是指针接受者定义的方法, 传入的person结构体是内存地址,所以readd中修改person的元素,就相当于修改person结构体的元素

}
