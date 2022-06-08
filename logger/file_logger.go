package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type FileLogger struct {
	FileName string
}


func (logger *FileLogger) Error(msg string) {
	writerFile(logger, ERROR.String(), msg)
}
func (logger *FileLogger) Info(msg string) {
	writerFile(logger, INFO.String(), msg)
}
func (logger *FileLogger) Debug(msg string) {
	writerFile(logger, DEBUG.String(), msg)
}

func checkFileIsExist(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

func writerFile(logger *FileLogger, level, msg string) {
	var file *os.File
	var err1 error
	if checkFileIsExist(logger.FileName) { //如果文件存在
		fmt.Println("=====>文件存在，进行追加！")
		file, err1 = os.OpenFile(logger.FileName, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
	} else {
		fmt.Println("=====>文件不存在,创建文件！")
		file, err1 = os.Create(logger.FileName) //创建文件
	}
	defer file.Close()
	var msgResult = time.Now().Format(time.RFC3339) + "\t" + level + "\t==>: " + msg
	n, err1 := io.WriteString(file, msgResult+"\n") //写入文件(字符串)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("=====>日志记录完毕，写入 %d 个字节\n", n)
}
