package main

import (
	"bufio"
	"fmt"
	proto "golang/08/ninabao/protocol"
	"io"
	"net"
)

// socket_stick/server/main.go
//客户端分10次发送的数据，在服务端并没有成功的输出10次，而是多条数据“粘”到了一起。因为发送tcp数据的时候先等一下,看有没有后续的数据,如果有会连在一起发送,这样就导致了黏包的发生
/*
主要原因就是tcp数据传递模式是流模式，在保持长连接的时候可以进行多次的收和发。

“粘包”可发生在发送端也可发生在接收端：

由Nagle算法造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法。简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，
若有则会一次把这两段数据发送出去。
接收端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层取数据。当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据
*/
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	// var buf [1024]byte
	for {
		// n, err := reader.Read(buf[:])
		// if err == io.EOF {
		// 	break
		// }
		// if err != nil {
		// 	fmt.Println("read from client failed, err:", err)
		// 	break
		// }
		// recvStr := string(buf[:n])
		recvStr, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println("收到client发来的数据：", recvStr)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for i := 0; i < 10; i++ {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
