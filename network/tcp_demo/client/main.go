package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端连接失败 ", err)
		return
	}
	fmt.Println("客户端连接成功")

	defer conn.Close()

	//从终端读取输入的数据
	reader := bufio.NewReader(os.Stdin)
	for true {

		linestr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("终端读取数据失败！ ", err)
			return
		}
		linestr = strings.TrimSpace(linestr)

		if strings.ToLower(linestr) == "exit" {
			fmt.Println("客户端退出了！ ")
			break
		}
		//把终端输入的字符串发送给服务端
		_, err = conn.Write([]byte(linestr))
		if err != nil {
			fmt.Println("客户端发送失败！ ", err)
			return
		}
	}
}
