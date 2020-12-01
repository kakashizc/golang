package main

import "fmt"

func main() {
	s1 := []int{1, 3, 5}
	s2 := s1
	s3 := make([]int, 5, 10)
	copy(s3, s1)
	s1[0] = 1111
	fmt.Printf("%v %v %v \n", s1, s2, s3)
	//[1111 2 3 4 5] [1111 2 3 4 5] [1 2 3 4 5]
	//因为s2是s1 赋值的, 所以内存地址一致 s1[0] 改变 , s2[0]也会变, 但是s3只是复制的s1的值, 内存地址跟s1不一样,所以s1改变,s3不变

	//切片的删除   切片是没有删除函数的,需要自己去做
	s1 = append(s1[:1], s1[2:]...) //从s1的第一位拼接 s1的第三个数 所以结果是 [1111 5]
	fmt.Println(s1)

	//切片底层分析
	x1 := [...]int{1, 3, 5}
	x2 := x1[:]            //x2是基于x1数组的一个切片
	fmt.Printf("%p\n", x2) //打印内存地址
	fmt.Printf("%p\n", &x2[0])
	x2 = append(x2[:1], x2[2:]...)
	fmt.Printf("%p\n", x2) //打印内存地址
	fmt.Printf("%p\n", &x2[0])
	fmt.Println(x2, len(x2), cap(x2)) //[1 5] 2 3 长度2,容量3
	//1,切片不保存具体的值
	//2,切片对应的就是一个底层数组
	//3,底层数组占用的是一块连续的内存
	fmt.Println(x1) // [1 5 5]  因为 x2切片的3 被删除了, 并且x2指向x1的内存地址, x1的3没有了,但是第二位被改成5了, 所以是1 5 5

	//底层再举例分析
	b1 := [...]int{1, 3, 5, 7, 9}
	s11 := b1[:]
	s11 = append(s11[:1], s11[2:]...)
	fmt.Println(s11, b1) //[1 5 7 9]  和 b1 = [1 5 7 9 9] 因为b1中的3被删除了,所以后面的579整体往前移动一位,最后的9仍然在内存块中保留,所以是15799

}
