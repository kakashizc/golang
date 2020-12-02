package calc //包名最好和上级文件名一致
import "fmt"

//初始化函数
func init() {
	fmt.Println("我被调用了")
}

//Add 首字母大写表示公开的方式,其他保中才能调用此方法
func Add(x, y int) int {
	return x + y
}
