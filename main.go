package main

import (
	"fmt"
	"io"
	"log"
)

// ANSI escape codes for text formatting
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
)

type LoggerConfig struct {
	WriteTo      io.Writer
	LogLevel     LogLevel
	WithDateTime bool
}

func New(c LoggerConfig) (*Logger, error) {
	var wr io.Writer = c.WriteTo
	var flag int
	if c.WithDateTime == true {
		flag = log.Ldate | log.Ltime | log.Lshortfile
	}
	InfoLogger := log.New(wr, fmt.Sprintf("%s[%s]%s ", Green, "INFO", Reset), flag)
	WarningLogger := log.New(wr, fmt.Sprintf("%s[%s]%s ", Yellow, "WARN", Reset), flag)
	ErrorLogger := log.New(wr, fmt.Sprintf("%s[%s]%s ", Red, "ERROR", Reset), flag)
	return &Logger{level: c.LogLevel, info: InfoLogger, warn: WarningLogger, error: ErrorLogger}, nil
}

// LogLevel represents different logging levels.
type LogLevel int

const (
	Error LogLevel = iota
	Warning
	Info
)

// Logger is a simple logger struct.
type Logger struct {
	level LogLevel
	info  *log.Logger
	error *log.Logger
	warn  *log.Logger
}

// Error logs an error message.
func (l *Logger) Error(message string) {
	if l.level <= Error {
		l.error.Println(message)
	}
}

// Warning logs a warning message.
func (l *Logger) Warning(message string) {
	if l.level <= Warning {
		l.warn.Println(message)
	}
}

// Info logs an info message.
func (l *Logger) Info(message string) {
	if l.level <= Info {
		l.info.Println(message)
	}
}
