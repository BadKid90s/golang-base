package main

import "goworkspase/logger"

func main() {
	var log logger.Logger
	log = &logger.ConsoleLogger{}
	log.Error("wo shi cuowu de rizhi")
	log.Info("wo shi cuowu de info")
	log.Debug("wo shi cuowu de debug")

	log = &logger.FileLogger{FileName: "logs.log"}
	log.Error("wo shi cuowu de rizhi")
	log.Info("wo shi cuowu de info")
	log.Debug("wo shi cuowu de debug")
}
