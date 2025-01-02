package logging

import (
	"io"
	"log/slog"

	"github.com/lmittmann/tint"
)

type Options struct {
	Debug  bool
	Writer io.Writer
}

var logger *slog.Logger

func Init(opts Options) {
	logLevel := slog.LevelInfo
	if opts.Debug {
		logLevel = slog.LevelDebug
	}
	logger = slog.New(tint.NewHandler(opts.Writer, &tint.Options{Level: logLevel}))
}

func Logger() *slog.Logger {
	if logger == nil {
		panic("logger is not initialized")
	}
	return logger
}
