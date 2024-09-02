package logger

import (
	"github.com/rs/zerolog"
	"io"
	"os"
	"strings"
	"time"
)

type Logger struct {
	Logger *zerolog.Logger
}

func New(level string) *Logger {
	var l zerolog.Level
	var loggerOutput io.Writer

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	loggerOutput = zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}

	if os.Getenv("ENV") != "development" {
		loggerOutput = os.Stderr
	}

	logger := zerolog.New(loggerOutput).With().Timestamp().Logger()

	return &Logger{
		Logger: &logger,
	}
}
