package config

import (
	"io"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

type Logger = log.Logger

func NewLogger(p string) *log.Logger {
	writer := io.Writer(os.Stdout)
	logger := log.NewWithOptions(writer, log.Options{
		Prefix:          p,
		Level:           log.DebugLevel,
		TimeFormat:      time.Kitchen,
		ReportCaller:    true,
		ReportTimestamp: true,
	})

	return logger
}

// //create Non-Formated Logs

// func (l *Logger) Debug(v ...interface{}) {
// 	l.debug.Println(v...)
// }

// func (l *Logger) Info(v ...interface{}) {
// 	l.info.Println(v...)
// }

// func (l *Logger) Warn(v ...interface{}) {
// 	l.warn.Println(v...)
// }

// func (l *Logger) Error(v ...interface{}) {
// 	l.err.Println(v...)
// }

// //create Formated enable Logs

// func (l *Logger) Debugf(format string, v ...interface{}) {
// 	l.debug.Printf(format, v...)
// }

// func (l *Logger) Infof(format string, v ...interface{}) {
// 	l.info.Printf(format, v...)
// }

// func (l *Logger) Warnf(format string, v ...interface{}) {
// 	l.warn.Printf(format, v...)
// }

// func (l *Logger) Errorf(format string, v ...interface{}) {
// 	l.err.Printf(format, v...)
// }
