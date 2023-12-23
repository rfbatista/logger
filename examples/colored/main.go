package main

import (
	"os"

	"github.com/rfbatista/logger"
)

func main() {
	c, _ := logger.New(logger.LoggerConfig{WriteTo: os.Stdout, LogLevel: logger.Debug, WithDateTime: true})
	c.Info("simple Info logging")
	c.Warning("simple Warning logging")
	c.Error("simple Error logging")
	c.Debug("simple Debug logging")
	c.Critical("simple Critical logging")
}
