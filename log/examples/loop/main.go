package main

// var config = log.Config{
// 	// zap
// 	Level:       "debug",
// 	Dev: false,

// 	Rotate: log.Rotate{
// 		Filename:   "biz.log",
// 		MaxSize:    100,
// 		MaxAge:     3600,
// 		MaxBackups: 5,
// 		LocalTime:  true,
// 		Compress:   true,
// 	},
// }

// var logger = log.New(config)

// func main() {

// 	log.SetDefault(logger)
// 	defer logger.Sync()

// 	for {
// 		// ctx := trace.Context(context.Background())
// 		// simpleHTTPGet(ctx, "http://baidu.com")
// 		time.Sleep(1 * time.Millisecond)
// 	}
// }

// func simpleHTTPGet(ctx context.Context, url string) {
// 	log.Info(ctx, "Trying to hit GET request for", zap.String("url:", url))
// 	get(ctx, url)
// 	simpleHTTPGet2(ctx, "http://alibaba.com")

// }

// func simpleHTTPGet2(ctx context.Context, url string) {
// 	log.Debug(ctx, "Trying2 to hit GET request for", zap.String("url:", url))
// 	get(ctx, url)
// 	simpleHTTPGet3(ctx, "http://tencent.com")
// }

// func simpleHTTPGet3(ctx context.Context, url string) {
// 	log.Debug(ctx, "Trying3 to hit GET request for", zap.String("url:", url))
// 	get(ctx, url)
// }

// func get(ctx context.Context, url string) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Error(ctx, "Error fetching", zap.String("url", url), zap.Error(err))
// 		return
// 	}
// 	resp.Body.Close()
// 	log.Info(ctx, "Success!", zap.String("url", url), zap.String("status", resp.Status))
// }
