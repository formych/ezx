package log

import (
	"github.com/fsm-xyz/ezx/config"
	log "github.com/fsm-xyz/ezx/log/zap"
	"github.com/fsm-xyz/ezx/log/zerolog"
)

func Init() {
	if config.C.Log.Provider == "zap" {
		log.Init()
	} else if config.C.Log.Provider == "zerolog" {
		zerolog.Init()
	}
	zerolog.Init()
}

func Close() {
	if config.C.Log.Provider == "zap" {
		log.Close()
	}
}
