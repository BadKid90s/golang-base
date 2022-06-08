package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 适用于较小的文件
	fileByte, err := ioutil.ReadFile("D:\\hello.txt")
	if err != nil {
		fmt.Println("文件打开错误", err)
	}
	fmt.Println(string(fileByte))
}
