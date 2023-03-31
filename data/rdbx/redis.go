package rdbx

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisMap = map[string]*redis.Client{}
)

func Init(c []Config) error {
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
	for _, v := range redisMap {
		if err := v.Close(); err != nil {
			// stdLogger.Error("close db failed", zap.String("name", k), zap.Error(err))
		}
	}
}

// Config 配置
// 同redis.Options
// 看情况暴露内部参数设置
type Config struct {
	Name     string `json:"name" yaml:"name"`
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
	// Maximum number of socket connections.
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	PoolSize int `json:"pool_size" yaml:"pool_size"`
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns int `json:"min_idle_conns" yaml:"min_idle_conns"`
	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration `json:"dial_timeout" yaml:"dial_timeout"`
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeout time.Duration `json:"read_timeout" yaml:"read_timeout"`
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeout time.Duration `json:"write_timeout" yaml:"write_timeout"`
}

// 后续增加Config配置验证，避免不合理的配置， Validition

// NewRedis ...
func New(r Config) *redis.Client {
	options := &redis.Options{
		Addr:         r.Addr,
		Password:     r.Password,
		PoolSize:     r.PoolSize,
		MinIdleConns: r.MinIdleConns,
		DB:           r.DB,
		DialTimeout:  r.DialTimeout,
		ReadTimeout:  r.ReadTimeout,
		WriteTimeout: r.WriteTimeout,
	}

	return redis.NewClient(options)
}
