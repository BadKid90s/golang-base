package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("D:\\hello123.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("文件打开错误", err)
		return
	}
	defer file.Close()
	str := "hello world \n"
	//写入时，使用带缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//刷新，真正写入到文件
	writer.Flush()
}
