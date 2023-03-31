package log

var Std *Logger

func Init(c Config) {
	// 默认的logger
	logger := New(c)
	SetDefault(logger)

	// 框架使用的std logger
	c.Outout = "std"
	Std = New(c)
}

func Close() {
	Sync()
	Std.Sync()
}
