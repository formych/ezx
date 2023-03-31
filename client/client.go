package client

import (
	"context"
	"fmt"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

var store = map[string]*grpc.ClientConn{}

// Config 配置  区分服务类型和地址
type Config struct {
	Name     string        `json:"name" yaml:"name"`
	Type     string        `json:"type" yaml:"type"`
	From     string        `json:"from" yaml:"from"`
	Endpoint string        `json:"endpoint" yaml:"endpoint"`
	Balancer string        `json:"balancer" yaml:"balancer"`
	Timeout  time.Duration `json:"timeout" yaml:"timeout"`
}

func (c Config) dial(addr string) (*grpc.ClientConn, error) {
	return c.direct()
}

// 直接进行连接
func (c Config) direct() (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	return grpc.DialContext(
		ctx,
		c.Endpoint,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
			// metrics.UnaryClientInterceptor,
			// trace.UnaryClientInterceptor,
			// cost.UnaryClientInterceptor,
			),
		),
	)
}

// ErrClientNotFound 未找到对应的client
var (
	ErrClientNotFound = "client not found, name: %s"
)

// client类型
const (
	grpcType = "grpc"
)

// Register 批量注册client
// 只支持grpc直连模式
func Register(addr string, clients []Config) {
	for _, c := range clients {
		if c.Type != grpcType {
			fmt.Printf("register client type not support, name: %s\n", c.Name)
			continue
		}

		conn, err := c.dial(addr)
		if err != nil {
			fmt.Printf("register client failed, name: %s, err: %s\n", c.Name, err)
			continue
		}
		fmt.Printf("register client success, name: %s\n", c.Name)
		store[c.Name] = conn
	}
}

// GetClient 提供外部Client进行获取grpc连接
func GetClient(name string) (*grpc.ClientConn, error) {
	if c, ok := store[name]; ok {
		return c, nil
	}
	return nil, fmt.Errorf(ErrClientNotFound, name)
}

// Close 关闭grpc连接
func Close() error {
	var err error
	for k, v := range store {
		if err = v.Close(); err != nil {
			err = fmt.Errorf("close grpc client, name: %s, err: %s", k, err)
		}
	}

	return err
}
