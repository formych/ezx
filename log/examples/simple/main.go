package main

// var config = log.Config{
// 	// zap
// 	Level:  "debug",
// 	Outout: "file",
// 	Dev:    false,

// 	Rotate: log.Rotate{
// 		Filename:   "example.log",
// 		MaxSize:    10,
// 		MaxAge:     3600,
// 		MaxBackups: 5,
// 		LocalTime:  true,
// 		Compress:   false,
// 	},
// }

// var stdConfig = log.Config{
// 	// zap
// 	Level:  "debug",
// 	Outout: "std",
// 	Dev:    false,
// }

// var glogger = log.New(config)

func main() {
	// defer glogger.Sync()
	// log.SetDefault(glogger)

	// // ctx := trace.Context(context.Background())
	// msg := "我是一条消息"
	// fileds := []zap.Field{zap.String("event", "test"), zap.String("addr", "北京/北京")}

	// // 使用log包级别的打印
	// // log.Info(ctx, msg, fileds...)

	// // 直接使用
	// glogger.Info(msg, fileds...)

	// // 默认字段
	// // logger := log.NewWith(ctx, zap.String("default", "default"))
	// logger.Info(msg, fileds...)

	// // // 控制台输出
	// var stdLogger = log.New(stdConfig)
	// stdLogger.Info(msg, fileds...)
}
