package log

import (
	"log"
	"os"
)

type ILogger interface {
	Info(string)
	Error(string)
}

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{log.New(os.Stderr, "payne: ", log.Lshortfile)}
}
