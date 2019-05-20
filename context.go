package logan

import (
	"context"
	"github.com/sirupsen/logrus"
)

type LogContext logrus.Fields

type key int

const logContextKey key = 0

func NewContextWithLogContext(ctx context.Context, logCtx LogContext) context.Context {
	return context.WithValue(ctx, logContextKey, logCtx)
}

func LogContextFromContext(ctx context.Context) (LogContext, bool) {
	logCtx, ok := ctx.Value(logContextKey).(LogContext)
	return logCtx, ok
}
