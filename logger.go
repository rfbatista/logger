package logger

import (
	"fmt"
	"io"
	"log"
)

// ANSI escape codes for text formatting
const (
	Reset    = "\033[0m"
	Red      = "\033[91m"
	Yellow   = "\033[93m"
	Green    = "\033[92m"
	Purple   = "\033[94m"
	BgRed    = "\033[41m"
	BgPurple = "\033[44m"
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
		flag = log.LUTC | log.Ldate | log.Ltime | log.Lshortfile
	}
	DebugLogger := log.New(wr, applyFormat(Purple, "DEBUG"), flag)
	InfoLogger := log.New(wr, applyFormat(Green, "INFO"), flag)
	WarningLogger := log.New(wr, applyFormat(Yellow, "WARN"), flag)
	ErrorLogger := log.New(wr, applyFormat(Red, "ERROR"), flag)
	CriticalLogger := log.New(wr, applyFormat(BgRed, "CRITICAL"), flag)
	if c.LogLevel > DebugLevel {
		DebugLogger = log.New(io.Discard, "", flag)
	}
	if c.LogLevel > InfoLevel {
		InfoLogger = log.New(io.Discard, "", flag)
	}
	if c.LogLevel > WarningLevel {
		WarningLogger = log.New(io.Discard, "", flag)
	}
	if c.LogLevel > ErrorLevel {
		ErrorLogger = log.New(io.Discard, "", flag)
	}
	if c.LogLevel > CriticalLevel {
		CriticalLogger = log.New(io.Discard, "", flag)
	}
	logger := &Logger{}
	logger.Debug = DebugLogger.Println
	logger.Info = InfoLogger.Println
	logger.Warning = WarningLogger.Println
	logger.Error = ErrorLogger.Println
	logger.Critical = CriticalLogger.Println
	return logger, nil
}

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	CriticalLevel
)

type Logger struct {
	Error    func(v ...any)
	Critical func(v ...any)
	Debug    func(v ...any)
	Info     func(v ...any)
	Warning  func(v ...any)
}
