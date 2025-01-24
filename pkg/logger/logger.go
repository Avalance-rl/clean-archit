package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

type Logger struct {
	*zap.Logger
}

var (
	logger *Logger
	once   sync.Once
)

// GetLogger возвращает единственный экземпляр Logger
func GetLogger() *Logger {
	once.Do(func() {
		logger = initLogger()
	})
	return logger
}

// initLogger инициализирует Logger с заданной конфигурацией
func initLogger() *Logger {
	config := zap.NewProductionConfig()

	config.OutputPaths = []string{
		"logs/application.log",
		"stderr",
	}
	config.ErrorOutputPaths = []string{
		"logs/error.log",
		"stderr",
	}

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	logLevel := os.Getenv("LOG_LEVEL")
	if level, err := zapcore.ParseLevel(logLevel); err == nil {
		config.Level = zap.NewAtomicLevelAt(level)
	}

	log, err := config.Build()
	if err != nil {
		panic("failed to build logger: " + err.Error())
	}

	return &Logger{Logger: log}
}
