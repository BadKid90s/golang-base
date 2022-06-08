package main

import (
	"encoding/json"
	"goworkspase/chartroom/define"
	"goworkspase/chartroom/helper"
	"goworkspase/chartroom/message"
	"log"
	"net"
)

func main() {
	listen, err := helper.CreateListen(define.ServerAddress)
	if err != nil {
		log.Printf("[创建监听错误：%v \n]", err.Error())
	}
	for true {
		tcpConn, err := listen.AcceptTCP()
		if err != nil {
			log.Printf("[建立连接错误：%v \n]", err.Error())
		}
		//启动协程，与客户端保持通信
		go process(tcpConn)
	}

}

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	//获取客户端传递的数据
	var buffs = make([]byte, 0)
	for true {
		var buf [1024]byte
		read, err := conn.Read(buf[:])
		if err != nil {
			return
		}
		buffs = append(buffs, buf[:read]...)
		if read < 1024 {
			break
		}
	}
	var mes message.Message
	err := json.Unmarshal(buffs, &mes)
	if err != nil {
		log.Printf("[反序列化信息错误：%v \n]", err.Error())

	}

}
