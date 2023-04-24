package log

import "github.com/fsm-xyz/ezx/config"

var Std *Logger

func Init() {
	// 默认的logger
	logger := New()
	SetDefault(logger)

	// 框架使用的std logger
	config.C.Log.Output = "stdout"
	Std = New()
}

func Close() {
	Sync()
	Std.Sync()
}
