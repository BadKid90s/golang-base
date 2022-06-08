package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//覆盖
	//file, err := os.OpenFile("D:\\hello123.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	// 追加
	//file, err := os.OpenFile("D:\\hello123.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//读写,追加
	file, err := os.OpenFile("D:\\hello123.txt", os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("文件打开错误", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for  {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	str := "hello world123456789bbb\r\n"
	//写入时，使用带缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//刷新，真正写入到文件
	writer.Flush()
}
