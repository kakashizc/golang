package main

import "fmt"

func main() {
	f1()
	f2(1, 2)
	a := f3()
	fmt.Println(a)
	b := f4(2, 3)
	fmt.Println(b)
	//接收多个返回值
	m, n := f5()
	fmt.Println(m, n)
	//如果只想要一个返回值
	_, k := f5()
	fmt.Println(k)

	//不定长参数
	f7("大家", 1, 3, 4)
}

//无参 无返回值
func f1() {
	fmt.Println("1")
}

//有参 无返回值
func f2(x int, y int) {
	fmt.Println(x + y)
}

//无参, 有返回值
func f3() (ret int) {
	ret = 1
	return
}

//有参数 有返回值
func f4(x int, y int) (ret int) {
	ret = x + y
	return
}

//多个返回值
func f5() (int, string) {
	return 1, "你好"
}

//参数的类型简写 当参数类型一致时可以简写
func f6(x, y int, m, n string) int {
	return x + y
}

//不定长参数.  y 可以传多个, 1,2,3
func f7(x string, args ...int) {
	fmt.Println(x)
	fmt.Println(args)
}
