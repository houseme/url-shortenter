package logger

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

// Logger is the interface that wraps the basic Logger methods.
type utilLogger struct {
	ctx    context.Context
	logger string
}

// Logger .
func Logger(ctx context.Context) *utilLogger {
	return &utilLogger{ctx: ctx}
}

// SetLogger set logger
func (l *utilLogger) SetLogger(ctx context.Context, logger string) context.Context {
	l.ctx = ctx
	l.logger = logger
	return context.WithValue(ctx, "logger", logger)
}

// GetLogger .get logger
func (l *utilLogger) GetLogger(ctx context.Context) string {
	return gconv.String(ctx.Value("logger"))
}
