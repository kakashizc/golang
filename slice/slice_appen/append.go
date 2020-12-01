package main

import "fmt"

func main() {
	s1 := []string{"shanghai", "beijing", "sjz"}
	//调用append函数, 必须要用原来的切片接收返回值
	s1 = append(s1, "guangzhou")
	//append 追加元素,原来的底层数组放不下了,
	//Go底层就会重新开一个容量更大的底层数组,所以必须要用一个值去接收,因为是追加所以我们用以前的值s1接收,当然也可以用其他接收
	fmt.Printf("%v", s1)
	//切片中追加切片
	s2 := []string{"shenzhen", "xiamen"}
	s1 = append(s1, s2...) //注意: 直接写s2 是不行的, 因为s2是切片类型,而s1中的元素是string类型,所以需要用 ... 把切片拆开后再追加
	fmt.Printf("%v", s1)

	//面试题
	a := make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a) //结果是: [0 0 0 0 0 1 2 3 4 5 6 7 8 9] 因为 make切片的时候 有5个长度了, 默认值是0

}
