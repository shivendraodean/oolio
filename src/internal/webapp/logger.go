package webapp

import (
	"log/slog"
	"os"
)

func InitLogger() *slog.Logger {
	options := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, options)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger
}
