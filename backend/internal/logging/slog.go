package logging

import (
	"io"
	"log/slog"

	"github.com/lmittmann/tint"
)

type Logger struct {
	*slog.Logger
	level *slog.LevelVar
}

func New(w io.Writer, level slog.Level) *Logger {
	levelVar := new(slog.LevelVar)
	levelVar.Set(level)
	return &Logger{
		Logger: slog.New(tint.NewHandler(w, &tint.Options{Level: levelVar})),
		level:  levelVar,
	}
}

func (logger *Logger) SetLevel(level slog.Level) {
	logger.level.Set(level)
}

func (logger *Logger) Level() slog.Level {
	return logger.level.Level()
}
