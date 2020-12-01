package main

import "fmt"

func main() {

	//切片的定义
	var s []int //定义了一个 没有固定元素个数的切片, 切片的底层就是数组
	fmt.Println(s)
	//make函数 切片
	s1 := make([]int, 5, 10) //声明一个int类型的切片, 长度是5,容量是10
	fmt.Printf("%d %d\n", len(s1), cap(s1))
	//切片的赋值 s2和 s3 指向同一个内存地址 所以s3的值改变,s2的值也会跟着改变
	s2 := []int{1, 3, 5, 7}
	s3 := s2
	fmt.Println(s2, s3)
	s3[0] = 1200
	fmt.Println(s2, s3)
}
