package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var a person
	a.name = "张三"
	a.age = 11
	fmt.Println(a)

	// 匿名结构体
	//声明方式一
	var p = person{
		name: "水电费",
		age:  12,
	}
	fmt.Println(p)
	//声明方式二
	p1 := person{
		"你好",
		11,
	}
	fmt.Println(p1)
}
