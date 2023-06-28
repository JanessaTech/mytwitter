package logging

import (
	"context"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/hi-supergirl/mytwitter/config"
	"go.uber.org/zap"
)

var once sync.Once
var logger *zap.Logger

const loggerKey = "logger"

func createLogger(isDevelopment bool) {
	once.Do(func() {
		var err error
		if isDevelopment {
			logger, err = zap.NewDevelopment()
		} else {
			logger, err = zap.NewProduction()
		}
		if err != nil {
			logger = zap.NewNop()
		}
		fmt.Println("logger is created")
	})
}

func GetLogger(config *config.Config) *zap.Logger {
	if logger == nil {
		createLogger(config.LoggingConfig.Development)
	}
	return logger
}

func DefaultLogger() *zap.SugaredLogger {
	if logger == nil {
		createLogger(true)
	}
	return logger.Sugar()
}

func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	if gCtx, ok := ctx.(*gin.Context); ok {
		ctx = gCtx.Request.Context()
	}
	return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	if gCtx, ok := ctx.(*gin.Context); ok {
		ctx = gCtx.Request.Context()
	}

	if logger, ok := ctx.Value(loggerKey).(*zap.SugaredLogger); ok {
		return logger
	}
	return DefaultLogger()
}
