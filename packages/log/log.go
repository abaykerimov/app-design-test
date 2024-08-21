package log

import (
	"fmt"
	"log"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	l := log.Default()
	return &Logger{l}
}

func (l Logger) LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.Printf("[Error]: %s\n", msg)
}

func (l Logger) LogInfoF(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.Printf("[Info]: %s\n", msg)
}

func (l Logger) LogFatalF(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.Fatalf("[Fatal]: %s\n", msg)
}
