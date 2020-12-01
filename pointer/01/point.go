package main

import "fmt"

func main() {
	//指针, go中的指针比较简单只能读取, 不像C语言可以修改指针
	//1, & 取内存地址

	a := 01
	fmt.Println(&a) //获取内存地址的值
	fmt.Println(a)
	p := &a // p的值就是一个内存地址 0xc000012090
	fmt.Println(p)
	fmt.Printf("%T", p) //*int  int类型的指针
	//2, * 根据地址取值
	m := *p        //去某个ip地址对应的值
	fmt.Println(m) // 1
}
