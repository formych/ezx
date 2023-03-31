package log

// Config 日志配置
type Config struct {
	// zap
	Level  string `json:"level" yaml:"level"`
	Outout string `json:"output" yaml:"output"`
	Dev    bool   `json:"dev" yaml:"dev"`

	// 日志切割配置
	// 当前采用 github.com/natefinch/lumberjack
	Rotate Rotate `json:"rotate" yaml:"rotate"`
}

// Rotate 日志切割配置
type Rotate struct {
	Filename   string `json:"filename" yaml:"filename"`
	MaxSize    int    `json:"maxsize" yaml:"maxsize"`
	MaxAge     int    `json:"maxage" yaml:"maxage"`
	MaxBackups int    `json:"maxbackups" yaml:"maxbackups"`
	LocalTime  bool   `json:"localtime" yaml:"localtime"`
	Compress   bool   `json:"compress" yaml:"compress"`
}
