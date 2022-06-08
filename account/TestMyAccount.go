package main

import "fmt"

func main() {
	var inputNum int
	var loop = true
	for {
		fmt.Println("------家庭收支记账软件------")
		fmt.Println("------1.收支明细 ------")
		fmt.Println("------2.登记收入 ------")
		fmt.Println("------3.登记支出 ------")
		fmt.Println("------4.退   出 ------")
		fmt.Print("		请选择1-4：")

		fmt.Scanln(&inputNum)

		switch inputNum {
		case 1:
			fmt.Println("------当前收支明细 ------")
		case 2:
			fmt.Println("------登记收入 ------")
		case 3:
			fmt.Println("------登记支出 ------")
		case 4:
			loop = false
		default:
			fmt.Println("------输入有误！ ------")
		}

		if !loop {
			break
		}
	}

	fmt.Println("推出系统！")

}
