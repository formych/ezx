package rdbx

import (
	"context"
	"fmt"

	"github.com/fsm-xyz/ezx/log"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	redisMap = map[string]*redis.Client{}
)

func Init(c []*Config) error {
	// 初始化redis
	for _, r := range c {
		client := New(r)
		if s, err := client.Ping(context.Background()).Result(); err != nil || s != "PONG" {
			return fmt.Errorf("name: %s, err:%s", r.Name, err)
		}
		redisMap[r.Name] = client
	}

	return nil
}

// GetRedis ...
func GetRedis(name string) *redis.Client {
	if v, ok := redisMap[name]; ok {
		return v
	}
	return nil
}

func Close() {
	// 关闭redis
	for k, v := range redisMap {
		if err := v.Close(); err != nil {
			log.Std.Error("close db failed", zap.String("name", k), zap.Error(err))
		}
	}
}

// 后续增加Config配置验证，避免不合理的配置， Validition

// NewRedis ...
func New(r *Config) *redis.Client {
	options := &redis.Options{
		Addr:         r.Addr,
		Password:     r.Password,
		PoolSize:     int(r.PoolSize),
		MinIdleConns: int(r.MinIdleConns),
		DB:           int(r.Db),
		DialTimeout:  r.DialTimeout.AsDuration(),
		ReadTimeout:  r.ReadTimeout.AsDuration(),
		WriteTimeout: r.WriteTimeout.AsDuration(),
	}

	return redis.NewClient(options)
}
