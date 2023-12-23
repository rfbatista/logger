package logger

import (
	"fmt"
	"io"
	"log"
)

// ANSI escape codes for text formatting
const (
	Reset  = "\033[0m"
	Red    = "\033[91m"
	Yellow = "\033[93m"
	Green  = "\033[92m"
	Purple  = "\033[94m"
	BgRed   = "\033[41m"
	BgPurple   = "\033[44m"
)

type LoggerConfig struct {
	WriteTo      io.Writer
	LogLevel     LogLevel
	WithDateTime bool
}

func applyFormat(color string, flag string) string {
	return fmt.Sprintf("%s[%s]%s ", color, flag, Reset)
}

func New(c LoggerConfig) (*Logger, error) {
	var wr io.Writer = c.WriteTo
	var flag int
	if c.WithDateTime == true {
		flag = log.Ldate | log.Ltime | log.Lshortfile
	}
	InfoLogger := log.New(wr, applyFormat(Green, "INFO"), flag)
	WarningLogger := log.New(wr, applyFormat(Yellow, "WARN"), flag)
	ErrorLogger := log.New(wr, applyFormat(Red, "ERROR"), flag)
	CriticalLogger := log.New(wr, applyFormat(BgRed, "CRITICAL"), flag)
	DebugLogger := log.New(wr, applyFormat(Purple, "DEBUG"), flag)
	return &Logger{level: c.LogLevel, info: InfoLogger, warn: WarningLogger, error: ErrorLogger, critical: CriticalLogger, debug: DebugLogger}, nil
}

type LogLevel int

const (
	Debug LogLevel = iota
	Error
	Warning
	Info
)

type Logger struct {
	level    LogLevel
	info     *log.Logger
	error    *log.Logger
	warn     *log.Logger
	critical *log.Logger
	debug    *log.Logger
}

func (l *Logger) Error(message string) {
	if l.level <= Error {
		l.error.Println(message)
	}
}

func (l *Logger) Warning(message string) {
	if l.level <= Warning {
		l.warn.Println(message)
	}
}

func (l *Logger) Info(message string) {
	if l.level <= Info {
		l.info.Println(message)
	}
}

func (l *Logger) Critical(message string) {
	l.critical.Println(message)
}

func (l *Logger) Debug(message string) {
	if l.level <= Debug {
		l.debug.Println(message)
	}
}
