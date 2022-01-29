package logger

import (
	"go.uber.org/zap"
	"log"
)

type Logger struct {
	zap   *zap.Logger
	sugar *zap.SugaredLogger
}

func New() *Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	return &Logger{
		zap:   logger,
		sugar: logger.Sugar(),
	}
}

func (l *Logger) Zap() *zap.Logger {
	return l.zap
}

func (l *Logger) Sugar() *zap.SugaredLogger {
	return l.sugar
}

func (l *Logger) ErrorWithStack(err error) {
	l.sugar.Fatal("fatal", zap.Error(err), zap.Stack("stack"))
}

func (l *Logger) Close() {
	if err := l.sugar.Sync(); err != nil {
		panic(err)
	}
	if err := l.zap.Sync(); err != nil {
		panic(err)
	}
}
