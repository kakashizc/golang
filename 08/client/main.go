package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//tcp 客户端
func main() {

	//1,与server端建立连接
	conn, err1 := net.Dial("tcp", "127.0.0.1:20000")
	if err1 != nil {
		fmt.Println("start tcp server failed,err:", err1)
		return
	}
	//2, 通信 发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		//1> 获取用户输入
		msg, _ := reader.ReadString('\n') //读到后换行
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		//2> 发送给服务端
		conn.Write([]byte(msg))
	}
	conn.Close()
}
