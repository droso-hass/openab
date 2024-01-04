package utils

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func SetupLogs(level string) {
	logLevel := new(slog.LevelVar)
	switch level {
	case "debug":
		logLevel.Set(slog.LevelDebug)
	case "info":
		logLevel.Set(slog.LevelInfo)
	case "warning":
		logLevel.Set(slog.LevelWarn)
	case "error":
		logLevel.Set(slog.LevelError)
	}
	handler := tint.NewHandler(os.Stdout, &tint.Options{
		Level:      logLevel,
		TimeFormat: time.Kitchen,
	})
	slog.SetDefault(slog.New(handler))
}
