package logger

import (
	"go.uber.org/zap"
	"log"
)

type Logger struct {
	*zap.Logger
}

func New() *Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	return &Logger{
		logger,
	}
}

func (l *Logger) ErrorWithStack(err error) {
	l.Fatal("fatal", zap.Error(err), zap.Stack("stack"))
}

func (l *Logger) Close() {
	if err := l.Sync(); err != nil {
		panic(err)
	}
}
