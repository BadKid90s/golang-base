package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("D:\\hello.txt")
	if err != nil {
		fmt.Println("文件打开错误", err)
	}
	//关闭文件
	defer file.Close()
	fmt.Printf("file=%v,type=%T \n", file, file)

	//创建一个reader,带缓冲的
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
}
