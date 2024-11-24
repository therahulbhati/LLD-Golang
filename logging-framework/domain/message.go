package domain

import "fmt"

type Level int

const (
	FATAL Level = iota
	ERROR
	WARN
	INFO
	DEBUG
)

type Message struct {
	Content   string
	Level     Level
	Namespace string
}

func (l Level) String() string {
	switch l {
	case FATAL:
		return "FATAL"
	case ERROR:
		return "ERROR"
	case WARN:
		return "WARN"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}

func ParseLevel(level string) (Level, error) {
	switch level {
	case "FATAL":
		return FATAL, nil
	case "ERROR":
		return ERROR, nil
	case "WARN":
		return WARN, nil
	case "INFO":
		return INFO, nil
	case "DEBUG":
		return DEBUG, nil
	default:
		return -1, fmt.Errorf("unknown level: %s", level)
	}
}
