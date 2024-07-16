package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	err    *log.Logger
	writer io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:  log.New(writer, "Debbug: ", logger.Flags()),
		info:   log.New(writer, "Info: ", logger.Flags()),
		warn:   log.New(writer, "Warning: ", logger.Flags()),
		err:    log.New(writer, "Error: ", logger.Flags()),
		writer: writer,
	}
}

//create Non-Formated Logs

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

//create Formated enable Logs

func (l *Logger) Debugf(format string,v ...interface{}) {
	l.debug.Printf(format,v...)
}

func (l *Logger) Infof(format string,v ...interface{}) {
	l.info.Printf(format,v...)
}

func (l *Logger) Warnf(format string,v ...interface{}) {
	l.warn.Printf(format,v...)
}

func (l *Logger) Errorf(format string,v ...interface{}) {
	l.err.Printf(format,v...)
}