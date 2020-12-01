package main

import "fmt"

func main() {
	//map类型, hash存储 键值对的类型
	m1 := make(map[string]int, 2)
	m1["张三"] = 11
	m1["李四"] = 12
	fmt.Println(m1)
	//判断map中是否存在某个键
	value, ok := m1["王五"]
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println(value)
	}

	//元素类型为map的切片
	s1 := make([]map[string]int, 5)
	s1[0] = make(map[string]int, 2) //s1切片中的map需要进行初始化
	s1[0]["北京"] = 111
	fmt.Println(s1)

	//值为切片的map
	m2 := make(map[string][]int, 10)
	m2["上海"] = []int{1, 3, 5}
	fmt.Println(m2)

}
