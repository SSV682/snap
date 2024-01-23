package logging

import (
	"context"

	"github.com/sirupsen/logrus"
)

// LoggerKey is a name of user field in context for storage logger key
const LoggerKey = "logger"

type Fields map[string]interface{}

var logger *logrus.Logger

// CtxWithLogger returns a copy of parent context with logger.
func CtxWithLogger(parentCtx context.Context, logger Entry) context.Context {
	return context.WithValue(parentCtx, LoggerKey, logger)
}

// GetLoggerFromContext get Entry from context.
// Notice: returns new Entry with the fields if empty value.
func GetLoggerFromContext(ctx context.Context) Entry {
	v := ctx.Value(LoggerKey)
	if l, ok := v.(Entry); ok {
		return l
	}

	return WithFields(make(Fields))
}

func WithFields(fields Fields) Entry {
	return NewEntry(logrus.NewEntry(logger)).WithFields(fields)
}
