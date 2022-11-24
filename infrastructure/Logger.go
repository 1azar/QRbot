package infrastructure

import (
	"log"
	"os"
)

type Logger struct {
	logInfo  *log.Logger
	logError *log.Logger
}

func (l Logger) Info(str ...any) {
	l.logInfo.Println(str)
}

func (l Logger) Error(str ...any) {
	l.logError.Println(str)
}

func NewLogger() Logger {
	return Logger{
		logInfo:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		logError: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
