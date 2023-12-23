package logger

import (
	"bytes"
	"fmt"
	"log"
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
		{"Error", Error, "Error message", fmt.Sprintf("%s[%s]%s Error message", Red, "ERROR", Reset)},
		{"Warning", Warning, "Warning message", fmt.Sprintf("%s[%s]%s Warning message", Yellow, "WARN", Reset)},
		{"Info", Info, "Info message", fmt.Sprintf("%s[%s]%s Info message", Green, "INFO", Reset)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect log output to a buffer for testing
			var buf bytes.Buffer
			log.SetOutput(&buf)

			logger, _ := New(LoggerConfig{LogLevel: tt.level, WithDateTime: false, WriteTo: &buf})

			switch tt.level {
			case Error:
				logger.Error(tt.message)
			case Warning:
				logger.Warning(tt.message)
			case Info:
				logger.Info(tt.message)
			}

			output := buf.String()

			// Check if the expected log message is present
			if strings.TrimRight(output, "\r\n") != tt.expected {
				t.Errorf("Expected log output: %s, got: %s", tt.expected, output)
			}

		})
	}
}
