package main

import (
	"fmt"
	proto "golang/08/ninabao/protocol"
	"net"
)

// socket_stick/client/main.go
//客户端分10次发送的数据，在服务端并没有成功的输出10次，而是多条数据“粘”到了一起。因为发送tcp数据的时候先等一下,看有没有后续的数据,如果有会连在一起发送,这样就导致了黏包的发生
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		//解决黏包 , 调用协议 编码数据
		m, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("error:", err)
		}
		conn.Write(m)
	}
}
