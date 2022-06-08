package logger

type LogLevel int16

const (
	ERROR LogLevel = 3
	INFO  LogLevel = 2
	WARN  LogLevel = 1
	DEBUG LogLevel = 0
)

func (l LogLevel) String() string {
	switch l {
	case ERROR:
		return "ERROR"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	}
	return "DEBUG"
}

