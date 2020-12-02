package main

import "fmt"

func main() {
	var x interface{}
	s := "石家庄"
	x = s
	fmt.Printf("类型是%T,%v", x, x)

	var stu = make(map[string]interface{}, 10)

	stu["name"] = "黄大"
	stu["age"] = 11
	stu["gender"] = true
	fmt.Println(stu)
}
