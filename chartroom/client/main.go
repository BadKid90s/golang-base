package main

import (
	"fmt"
	"goworkspase/chartroom/login"
)

var userId int

var userPassword string

func main() {

	//用户的输入
	var key int
	//循环
	loop := true

	for loop {
		fmt.Println("---------------聊天室-----------------")
		fmt.Println("\t  1.登录聊天室")
		fmt.Println("\t  2.注册用户")
		fmt.Println("\t  3.推出系统")
		fmt.Println("\t  请选择（1-3）：")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			println("推出系统")
			loop = false
		default:
			fmt.Println("输入有误，重新选择")
		}
	}

	//增加用户的信息
	if key == 1 {
		fmt.Println("请输入用户Id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%d\n", &userPassword)
		//登录
		err := login.Login(userId, userPassword)
		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}
	} else {

	}
}
