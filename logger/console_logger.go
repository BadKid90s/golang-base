package logger

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
	FileName string
}

func (logger *ConsoleLogger) Error(msg string) {
	writerConsole(ERROR.String(), msg)
}
func (logger *ConsoleLogger) Info(msg string) {
	writerConsole(INFO.String(), msg)
}
func (logger *ConsoleLogger) Debug(msg string) {
	writerConsole(DEBUG.String(), msg)
}

func writerConsole(level, msg string) {
	var msgResult = time.Now().Format(time.RFC3339) + "\t" + level + "\t==>: " + msg
	fmt.Println(msgResult)
}
