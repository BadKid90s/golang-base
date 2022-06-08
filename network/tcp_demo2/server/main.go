package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

var session = make(map[string]net.Conn)

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
			//存放连接消息
			split := strings.Split(remoteAddr.String(), ":")
			session[split[1]] = conn
			go process(conn)
		}
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for true {
		var buffs []byte
		for true {
			var buf [5]byte
			read, err := conn.Read(buf[:])
			if err != nil {
				fmt.Printf("读取数据出错了！%v", err.Error())
			}
			buffs = append(buffs, buf[:read]...)
			if read < 5 {
				break
			}
		}
		str := string(buffs[:])

		fmt.Printf("服务端接受到的消息：%v\n", str)

		if strings.HasPrefix(str, "conn:") {
			targetIpAddr := strings.ReplaceAll(str, "conn:", "")
			fmt.Printf("转发请求给客户端: %v\n", targetIpAddr)
			if targetConn, ok := session[targetIpAddr]; ok {
				fmt.Printf("进行转发！ \n")
				go io.Copy(conn, targetConn)
				go io.Copy(targetConn, conn)
			} else {
				fmt.Printf("转发客户端不存在: \n")
			}
		}

		conn.Write([]byte("服务端收到！"))

	}
}
