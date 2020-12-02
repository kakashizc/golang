package main

import "fmt"

//接口断言

func assign(a interface{}) {
	fmt.Printf("参数类型是:%T,值是%v", a, a)
	v, ok := a.(int) //返回值是ok, 如果a.(string) 返回true 那么ok=true
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Printf("答对了,类型就是%T", v)
	}

}

func main() {
	assign(100)
}
