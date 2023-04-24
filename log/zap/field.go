package log

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// getTrace 获取trace
func getTrace(ctx context.Context) zapcore.Field {
	return zap.Skip()
	// if ctx == nil {
	// 	return zap.Skip()
	// }

	// return zap.String("trace", trace.Trace(ctx))
}
