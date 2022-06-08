package main

import (
	"fmt"
	"net"
)

func main() {
	//创建监听
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听失败：", err)
		return
	}
	//延时关闭
	defer listener.Close()

	//循环等待客户端连接
	for true {
		fmt.Println("等待客户端连接。。。。")

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("创建连接失败：", err)
		} else {
			remoteAddr := conn.RemoteAddr()
			fmt.Println("IP：", remoteAddr, " 连接成功！")
			go process(conn)
		}
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for true {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("接受消息错误", err)
			continue
		}
		str := string(buf[:n])
		fmt.Println("服务端接受到的消息：", str)

	}
}
