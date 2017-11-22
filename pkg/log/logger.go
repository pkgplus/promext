package log

import (
	"fmt"
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, args ...interface{})
	Error(v ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, args ...interface{})
	Info(v ...interface{})
	Infof(format string, args ...interface{})
	Print(v ...interface{})
	Println(v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, args ...interface{})
}

type StdoutLogger struct{}

func (logger *StdoutLogger) Debug(v ...interface{})                    { logger.Print(v...) }
func (logger *StdoutLogger) Debugf(format string, args ...interface{}) { logger.Printf(format, args...) }
func (logger *StdoutLogger) Error(v ...interface{})                    { logger.Print(v...) }
func (logger *StdoutLogger) Errorf(format string, args ...interface{}) { logger.Printf(format, args...) }
func (logger *StdoutLogger) Fatal(v ...interface{})                    { logger.Print(v...) }
func (logger *StdoutLogger) Fatalf(format string, args ...interface{}) { logger.Printf(format, args...) }
func (logger *StdoutLogger) Info(v ...interface{})                     { logger.Print(v...) }
func (logger *StdoutLogger) Infof(format string, args ...interface{})  { logger.Printf(format, args...) }
func (logger *StdoutLogger) Warn(v ...interface{})                     { logger.Print(v...) }
func (logger *StdoutLogger) Warnf(format string, args ...interface{})  { logger.Printf(format, args...) }
func (logger *StdoutLogger) Print(v ...interface{})                    { fmt.Print(v...) }
func (logger *StdoutLogger) Println(v ...interface{})                  { fmt.Print(v...) }
func (logger *StdoutLogger) Printf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
