package log

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger ...
type Logger struct {
	logger *zap.Logger
}

// New 返回一个Logger
func New(config Config) *Logger {
	writer := getLogWriter(config)
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writer, getLevel(config.Level))

	if config.Dev {
		return &Logger{
			logger: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development()),
		}
	}
	return &Logger{logger: zap.New(core)}
}

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// 返回指定的level， 默认info
func getLevel(level string) zapcore.Level {
	if x, ok := levelMap[level]; ok {
		return x
	}
	return zap.InfoLevel
}

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	config.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewJSONEncoder(config)
}

// 日志输出位置
const (
	stdoutOutput = "stdout"
	stderrOutput = "stderr"
	fileOutput   = "file"
)

func getLogWriter(config Config) zapcore.WriteSyncer {

	if config.Outout == stdoutOutput {
		return os.Stdout
	}
	if config.Outout == stderrOutput {
		return os.Stderr
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.Rotate.Filename,
		MaxSize:    config.Rotate.MaxSize,
		MaxAge:     config.Rotate.MaxAge,
		MaxBackups: config.Rotate.MaxBackups,
		LocalTime:  config.Rotate.LocalTime,
		Compress:   config.Rotate.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// Debug ...
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

// Info ...
func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

// Warn ...
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

// Error ...
func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

// DPanic ...
func (l *Logger) DPanic(msg string, fields ...zap.Field) {
	l.logger.DPanic(msg, fields...)
}

// Panic ...
func (l *Logger) Panic(msg string, fields ...zap.Field) {
	l.logger.Panic(msg, fields...)
}

// Fatal ...
func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

// Sync ...
func (l *Logger) Sync() error {
	return l.logger.Sync()
}

// 包级别的logger
var (
	defaultL *Logger
)

// SetDefault 设置默认
func SetDefault(logger *Logger) {
	defaultL = logger
}

// GetLogger 提供默认logger, 实现自定义修改
func GetLogger() *Logger {
	return defaultL
}

// NewWith ...
func NewWith(ctx context.Context, fields ...zap.Field) *Logger {
	fields = append(fields, getTrace(ctx))
	return &Logger{logger: defaultL.logger.With(fields...)}
}

// Debug ...
func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getTrace(ctx))
	defaultL.logger.Debug(msg, fields...)
}

// Info ...
func Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getTrace(ctx))
	defaultL.logger.Info(msg, fields...)
}

// Warn ...
func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getTrace(ctx))
	defaultL.logger.Warn(msg, fields...)
}

// Error ...
func Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getTrace(ctx))
	defaultL.logger.Error(msg, fields...)
}

// DPanic ...
func DPanic(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getTrace(ctx))
	defaultL.logger.DPanic(msg, fields...)
}

// Panic ...
func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getTrace(ctx))
	defaultL.logger.Panic(msg, fields...)
}

// Fatal ...
func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getTrace(ctx))
	defaultL.logger.Fatal(msg, fields...)
}

// Sync ...
func Sync() error {
	return defaultL.logger.Sync()
}

// NewDataLog 创建只用来保存数据的logger
func NewDataLog(filename string) *Logger {
	var c = Config{
		Level:  "info",
		Outout: "file",
		Dev:    false,

		Rotate: Rotate{
			Filename:   filename,
			MaxSize:    1 << 6,
			MaxAge:     3,
			MaxBackups: 1,
			LocalTime:  true,
			Compress:   false,
		},
	}

	return New(c)
}
