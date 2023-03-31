# log

提供一个通用的logger组件

## 特性

+ trace打印
+ 携带默认字段
+ 支持控制台和文件输出

## 使用方式

+ 包级别
+ 对象方式
+ 新对象携带默认字段

```go

var glogger = log.New(config)

func main() {
    defer glogger.Sync()
	
	ctx := context.Background()
	msg := "我是一条消息"
	fileds := []zap.Field{zap.String("event", "test"), zap.String("addr", "北京/北京")}

	// 包级别的日志函数传入ctx即可打印日志
	log.Info(ctx, msg, fileds...)

	// 对象方式
	glogger.Info(msg, fileds...)

    // 设置默认的打印字段，自动加上trace信息
	logger := log.NewWith(ctx, zap.String("default", "default"))
    logger.Info(msg, fileds...)
    
    // 控制台输出
    log.StdL.Info(msg, fileds)
```

## Config

```go

var config = log.Config{
	// zap
    Level:       "debug",
    // 日志输出位置
    // 默认std, file为文件输出
	Outout:      "file",
	Development: false,

	// 日志切割配置 github.com/natefinch/lumberjack
	Rotate: log.Rotate{
		Filename:   "example.log",
		MaxSize:    100,
		MaxAge:     3600,
		MaxBackups: 5,
		LocalTime:  true,
		Compress:   false,
	},
}
```
