package ezx

import (
	"fmt"

	"github.com/fsm-xyz/ezx/config"
	"github.com/fsm-xyz/ezx/data/dbx"
	"github.com/fsm-xyz/ezx/data/rdbx"
	"github.com/fsm-xyz/ezx/log"
	"github.com/fsm-xyz/ezx/micro/kratos"
	"go.uber.org/zap"
)

// Engine ...
type Engine struct{}

// New ...
func New(bc any) (e *Engine) {
	// 初始化配置
	config.Init(bc)
	// 初始化logger
	log.Init(config.C.Log)

	// 打印配置
	// stdLogger.Info("config info", zap.ByteString("service", gdata), zap.ByteString("custem", cdata))

	// db资源初始化
	if err := dbx.Init(config.C.Data.Db); err != nil {
		log.Std.Fatal("init db failed", zap.Error(err))
		return
	}

	// db资源初始化
	if err := rdbx.Init(config.C.Data.Redis); err != nil {
		log.Std.Fatal("init redis failed", zap.Error(err))
		return
	}
	// 初始化成功
	// initMonitor(config.C.Monitor)
	return &Engine{}

}
func (e *Engine) Run() error {
	// 关闭资源
	defer Close()
	// 处理业务自定义的退出逻辑
	defer runHandlers()

	if config.C.Server.Provider == "kratos" {
		if err := kratos.Run(); err != nil {
			log.Std.Error("service exit", zap.Error(err))
			return err
		}
	}

	return nil
}

// 业务自定义的退出逻辑
var handlers = []func(){}

func runHandler(handler func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Std.Error("exit handler error:", zap.Error(fmt.Errorf("%v", err)))
		}
	}()

	handler()
}

func runHandlers() {
	for _, handler := range handlers {
		runHandler(handler)
	}
}

// RegisterExitHandlers 提供业务注册退出函数
func RegisterExitHandlers(hs ...func()) {
	handlers = append(handlers, hs...)
}

// Close 关闭资源
func Close() {
	dbx.Close()
	rdbx.Close()
	log.Close()
}
