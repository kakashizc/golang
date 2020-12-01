package main

import "fmt"

//构造函数
type person struct {
	name string
	age  int
}

//构造函数 方式一: 复制一份person, 如果person字段很多的话, 比较占用内存
// func newPerson(name string, age int) person {
// 	return person{
// 		name,
// 		age,
// 	}
// }

//方式二: 直接复制一份person的内存地址,地址就很小了,这样省内存
func newPerson(name string, age int) *person {
	return &person{
		name,
		age,
	}
}

func main() {
	a := newPerson("张三", 11)
	fmt.Println(a.name)
}
