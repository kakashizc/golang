package main

import "fmt"

func main() {
	//数组类型
	var arr [3]bool
	fmt.Printf("%T\n", arr)
	//初始化方式1
	arr = [3]bool{true, true, true}
	//初始化方式2,不知道数组的个数,可以用...来代替
	arr2 := [...]int{1, 2, 3, 4}
	fmt.Printf("%T\n", arr2)
	//初始化方式3, 索引方式:0:1 第一位的值是1  4:2,第5位的值是2
	arr3 := [5]int{0: 1, 4: 2} //数组的值: [1 0 0 0 2]
	fmt.Println(arr3)
	//循环数组
	citys := [...]string{"北京", "上海", "广州"}
	for k, v := range citys {
		fmt.Println(k, v)
	}
	//多维数组 [3] 代表第一维数组有3个元素,[2]代表每个二维数组有两个元素
	a4 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a4, "\n----------------")
	//循环二维数组
	for _, v1 := range a4 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//练习题: 1,求数组 [1 3 5 7 9]所有元素的和
	//		  2, 找出数组中 两个元素相加等于8的元素的下标
	arr5 := [...]int{1, 3, 5, 7, 9}
	var sum int
	for _, v1 := range arr5 {
		sum += v1
	}
	fmt.Println(sum)

	for i, v1 := range arr5 {
		for j, v2 := range arr5 {
			if v1+v2 == 8 {
				fmt.Println(i, j)
			}
		}
	}
}
