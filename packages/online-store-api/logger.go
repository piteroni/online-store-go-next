package common

import (
	"fmt"
	"io"
	"log"
	"runtime"
)

type AppLogger struct {
	logger *log.Logger
}

func NewLogger(f io.Writer) *AppLogger {
	logger := log.New(f, "", log.LstdFlags|log.Ldate|log.Lmsgprefix)

	return &AppLogger{
		logger: logger,
	}
}

func (l *AppLogger) Print(message string) {
	l.setCaller()

	l.logger.Println(message)
}

func (l *AppLogger) Printf(format string, v ...interface{}) {
	l.setCaller()

	l.logger.Printf(format, v...)
}

func (l *AppLogger) Error(err error) {
	l.setCaller()

	l.logger.Printf("%+v\n", err)
}

func (l *AppLogger) setCaller() {
	_, file, line, _ := runtime.Caller(2)

	l.logger.SetPrefix(fmt.Sprintf("%s:%d ", file, line))
}
