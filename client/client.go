package client

import (
	"context"
	"fmt"

	"github.com/fsm-xyz/ezx/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var store = map[string]*grpc.ClientConn{}

// Client 配置  区分服务类型和地址
type Client struct {
	*config.Client
}

func (c *Client) dial() (*grpc.ClientConn, error) {
	return c.direct()
}

// 直接进行连接
func (c *Client) direct() (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout.AsDuration())
	defer cancel()
	return grpc.DialContext(
		ctx,
		c.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		// grpc.WithUnaryInterceptor(
		// grpc_middleware.ChainUnaryClient(
		// metrics.UnaryClientInterceptor,
		// trace.UnaryClientInterceptor,
		// cost.UnaryClientInterceptor,
		// ),
		// ),
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
func Register(clients []*config.Client) {
	for k := range clients {
		c := &Client{
			clients[k],
		}
		if c.Type != grpcType {
			fmt.Printf("register client type not support, name: %s\n", clients[k].Name)
			continue
		}

		conn, err := c.dial()
		if err != nil {
			fmt.Printf("register client failed, name: %s, err: %s\n", clients[k].Name, err)
			continue
		}
		fmt.Printf("register client success, name: %s\n", clients[k].Name)
		store[clients[k].Name] = conn
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
