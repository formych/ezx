package zerolog

import (
	"io"
	"os"

	"github.com/fsm-xyz/ezx/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() {
	zerolog.LevelFieldName = "l"
	zerolog.ErrorFieldName = "e"
	zerolog.MessageFieldName = "m"
	zerolog.TimestampFieldName = "t"
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.000"

	c := config.C.Log
	writer := getLogWriter()

	if c.Dev {
		log.Logger = log.With().Logger()
		log.Logger = zerolog.New(writer).With().Caller().Timestamp().Logger()
	} else {
		log.Logger = zerolog.New(writer).With().Timestamp().Logger()
	}
}

// 日志输出位置
const (
	stdoutOutput = "stdout"
	stderrOutput = "stderr"
	fileOutput   = "file"
)

func getLogWriter() io.Writer {
	c := config.C.Log
	if c.Output == stdoutOutput {
		return os.Stdout
	}
	if c.Output == stderrOutput {
		return os.Stderr
	}

	return &lumberjack.Logger{
		Filename:   c.Rotate.Filename,
		MaxSize:    int(c.Rotate.MaxSize),
		MaxAge:     int(c.Rotate.MaxAge),
		MaxBackups: int(c.Rotate.MaxBackups),
		LocalTime:  c.Rotate.LocalTime,
		Compress:   c.Rotate.Compress,
	}
}
