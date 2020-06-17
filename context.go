package logger

import (
	"context"

	"github.com/qumonintelligence/go-logger/v2/console"
)

// loggerContextKey for logger
type loggerContextKey string

const defaultLoggerContext = loggerContextKey("LOGGER")

// FromContext get a logger from context
func FromContext(ctx context.Context) ILogger {
	value := ctx.Value(defaultLoggerContext)
	if value == nil {
		return console.GetLogger(defaultLogger)
	}

	log, ok := value.(ILogger)
	if !ok {
		return console.GetLogger(defaultLogger)
	}

	return log
}

// SaveContext save the logger to default context
func SaveContext(ctx context.Context, logger ILogger) context.Context {
	return context.WithValue(ctx, defaultLoggerContext, logger)
}
