package logdiscard

import (
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"github.com/rs/zerolog"
	"io"
)

func NewDiscardLogger() *logger.Logger {
	zl := zerolog.New(io.Discard)

	return &logger.Logger{Logger: &zl}
}
