package main

import "fmt"

//指针接口

//Mover is a type
type Mover interface {
	move()
}

type dog struct{}

func (d *dog) move() {
	fmt.Println("狗会动")
}
func main() {
	var x Mover        // x可以接收dog类型
	var fugui = &dog{} // 富贵是*dog类型
	x = fugui          // x可以接收*dog类型
	x.move()
}
