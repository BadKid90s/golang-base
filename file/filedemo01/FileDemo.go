package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("D:\\hello.txt")
	if err != nil {
		fmt.Println("文件打开错误", err)
	}
	fmt.Printf("file=%v,type=%T", file, file)
	//关闭文件
	defer file.Close()
}
