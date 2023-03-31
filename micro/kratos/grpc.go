package kratos

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"xyz-zyx.co/config"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *config.Server) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Network != "" {
		opts = append(opts, grpc.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, grpc.Address(c.Addr))
	}
	if c.Timeout > 0 {
		opts = append(opts, grpc.Timeout(c.Timeout))
	}
	srv := grpc.NewServer(opts...)
	return srv
}
