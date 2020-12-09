package main

import (
	"fmt"
	"net"
)

//TCP 客户端
func main() {
	//1,本地端口启动
	listener, err1 := net.Listen("tcp", "127.0.0.1:20000")
	if err1 != nil {
		fmt.Println("start tcp server failed,err:", err1)
		return
	}
	//2,等待别人跟我建立连接
	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("accept tcp server failed,err:", err2)
			return
		}
		go process(conn)
	}
}

//参数是net 包里面的 Conn 类型
func process(conn net.Conn) {
	//3,与客户端通信
	var temp [128]byte //定义一个字节类型的切片 Read() 参数需要
	for {
		n, err3 := conn.Read(temp[:])
		if err3 != nil {
			fmt.Println("Read failed,err:", err3)
			return
		}
		fmt.Println(string(temp[:n]))
	}
}
