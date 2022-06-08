package login

import (
	"encoding/json"
	"goworkspase/chartroom/define"
	"goworkspase/chartroom/helper"
	"goworkspase/chartroom/message"
	"log"
)

//Login 登录
func Login(userId int, userPassword string) (err error) {
	connect, err := helper.CreateConnect(define.ServerAddress)
	if err != nil {
		log.Printf("[建立连接错误：%v \n]", err.Error())
		return err
	}
	defer connect.Close()
	//实际消息
	var mes message.Message
	mes.Type = message.LoginMesType
	mes.Data = message.LoginMes{
		UserId:       userId,
		UserPassword: userPassword,
	}

	bytes, err := json.Marshal(mes)
	if err != nil {
		log.Printf("[序列化消息错误：%v \n]", err.Error())
		return err
	}

	_, err = connect.Write(bytes)
	if err != nil {
		log.Printf("[消息发送错误：%v \n]", err.Error())
		return err
	}
	return nil
}
