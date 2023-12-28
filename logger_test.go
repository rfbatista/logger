package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	tests := []struct {
		name     string
		level    LogLevel
		message  string
		expected string
	}{
		{"Error", ErrorLevel, "Error message", fmt.Sprintf("%s[%s]%s Error message", Red, "ERROR", Reset)},
		{"Warning", WarningLevel, "Warning message", fmt.Sprintf("%s[%s]%s Warning message", Yellow, "WARN", Reset)},
		{"Info", InfoLevel, "Info message", fmt.Sprintf("%s[%s]%s Info message", Green, "INFO", Reset)},
		{"Debug", DebugLevel, "Debug message", fmt.Sprintf("%s[%s]%s Debug message", Purple, "DEBUG", Reset)},
		{"Critical", CriticalLevel, "Critical message", fmt.Sprintf("%s[%s]%s Critical message", BgRed, "CRITICAL", Reset)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			logger, _ := New(LoggerConfig{LogLevel: tt.level, WithDateTime: false, WriteTo: &buf})

			switch tt.level {
			case ErrorLevel:
				logger.Error(tt.message)
			case WarningLevel:
				logger.Warning(tt.message)
			case InfoLevel:
				logger.Info(tt.message)
			case CriticalLevel:
				logger.Critical(tt.message)
			case DebugLevel:
				logger.Debug(tt.message)
			}

			output := buf.String()

			if strings.TrimRight(output, "\r\n") != tt.expected {
				t.Errorf("Expected log output: %s, got: %s", tt.expected, output)
			}

		})
	}

	t.Run("should not log info message if level is error", func(t *testing.T) {
		var buf bytes.Buffer
		logger, _ := New(LoggerConfig{LogLevel: ErrorLevel, WithDateTime: false, WriteTo: &buf})
		logger.Info("info message")
		output := buf.String()
		if output != "" {
			t.Errorf("expected to not log info message : %s", output)
		}
	})

	t.Run("should not log debug message if level is info", func(t *testing.T) {
		var buf bytes.Buffer
		logger, _ := New(LoggerConfig{LogLevel: InfoLevel, WithDateTime: false, WriteTo: &buf})
		logger.Debug("info message")
		output := buf.String()
		if output != "" {
			t.Errorf("expected to not log info message : %s", output)
		}
	})

	t.Run("should not log warning message if level is error", func(t *testing.T) {
		var buf bytes.Buffer
		logger, _ := New(LoggerConfig{LogLevel: ErrorLevel, WithDateTime: false, WriteTo: &buf})
		logger.Warning("info message")
		output := buf.String()
		if output != "" {
			t.Errorf("expected to not log info message : %s", output)
		}
	})


}
