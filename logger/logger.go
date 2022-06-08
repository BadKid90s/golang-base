package logger

type Logger interface {
	// Error 错误
	Error(msg string)
	// Info 信息
	Info(msg string)
	// Debug 错误
	Debug(msg string)
}
