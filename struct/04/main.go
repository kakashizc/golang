package main

import (
	"fmt"
	"os"
)

//练习: 学生管理系统,完成 -> 增删查功能

//一,首先需要做一个学生的结构体
type student struct {
	id   int
	name string
}

//二,需要一个变量去保存学生 用切片也可以,用map也行
var stu = make(map[int]*student, 50)

//学生的构造函数 构造函数new开头 用*指针 避免普通方式复制多个占用内存
func newStudent(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func main() {

	//用for循环是为了用户输入后,执行完继续来显示选项. 如果选择了4退出就直接退出for循环
	for {

		//1,展示选项
		fmt.Println("学生管理系统")
		fmt.Println(`
		1,查看所有学生
		2,新增学生
		3,删除学生
		4,退出
	`)
		//2,获取用户输入
		fmt.Println("请输入你的选项:")
		var inputNum int
		fmt.Scanln(&inputNum)
		fmt.Println("您选择了->", inputNum)
		//3,执行输入对应的函数
		switch inputNum {
		case 1:
			showStu()
		case 2:
			addStu()
		case 3:
			delStu()
		case 4:
			os.Exit(1) //退出
		default:
			fmt.Println("滚~")
		}
	}
}

//查看所有学生
func showStu() {
	for k, v := range stu {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}

//新增学生
func addStu() {
	//获取用户输入的值
	var (
		id   int
		name string
	)
	fmt.Println("请输入学号:")
	fmt.Scanln(&id)
	fmt.Println("请输入姓名:")
	fmt.Scanln(&name)
	//1,创建一个学生 方式一:
	// student := student{
	// 	id,
	// 	name,
	// }
	//1,创建一个学生, 方拾二: 用构造函数方法
	student := newStudent(id, name)
	//2,添加到stu这个map变量中
	stu[id] = student
}

//删除学生
func delStu() {
	var id int
	fmt.Println("请输入要删除的学生学号:")
	fmt.Scanln(&id)
	//去 stu 这个map中查找学号并删除
	delete(stu, id)
}
