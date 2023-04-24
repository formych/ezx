package ezx

import (
	"encoding/json"
	"fmt"

	"github.com/fsm-xyz/ezx/config"
	"github.com/fsm-xyz/ezx/data/dbx"
	"github.com/fsm-xyz/ezx/data/rdbx"
	"github.com/fsm-xyz/ezx/log"
	"github.com/fsm-xyz/ezx/micro/kratos"
	"google.golang.org/protobuf/encoding/protojson"

	zlog "github.com/rs/zerolog/log"
)

// Engine ...
type Engine struct{}

// New ...
func New(bc any) (e *Engine) {
	// 初始化配置
	config.Init(bc)
	// 初始化logger
	log.Init()

	// 打印配置
	printConfig(bc)
	// metircs
	go metircs()

	if config.C.Server.Provider == "kratos" {
		kratos.Init()
	}

	// db资源初始化
	if err := dbx.Init(config.C.Data.Db); err != nil {
		zlog.Fatal().Err(err).Msg("init db failed")
		return
	}

	// db资源初始化
	if err := rdbx.Init(config.C.Data.Redis); err != nil {
		zlog.Fatal().Err(err).Msg("init redis failed")
		return
	}

	return &Engine{}

}
func printConfig(bc any) {
	data, _ := protojson.Marshal(&config.C)
	bdata, _ := json.Marshal(&bc)
	zlog.Info().Bytes("data", data).Bytes("bdata", bdata).Msg("config info")
}
func (e *Engine) Run() error {
	// 关闭资源
	defer e.Close()
	// 处理业务自定义的退出逻辑
	defer e.runHandlers()

	if config.C.Server.Provider == "kratos" {
		if err := kratos.Run(); err != nil {
			zlog.Error().Err(err).Msg("service exit")
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
			zlog.Error().Err(fmt.Errorf("%s", err)).Msg("exit handler error")
		}
	}()

	handler()
}

func (e *Engine) runHandlers() {
	for _, handler := range handlers {
		runHandler(handler)
	}
}

// RegisterExitHandlers 提供业务注册退出函数
func (e *Engine) RegisterExitHandlers(hs ...func()) {
	handlers = append(handlers, hs...)
}

// Close 关闭资源
func (e *Engine) Close() {
	dbx.Close()
	rdbx.Close()
	log.Close()
}
