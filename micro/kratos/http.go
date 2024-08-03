package kratos

import (
	"github.com/fsm-xyz/ezx/config"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/durationpb"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *config.Service_Server) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Network != "" {
		opts = append(opts, http.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, http.Address(c.Addr))
	}
	if c.Timeout != durationpb.New(0) {
		opts = append(opts, http.Timeout(c.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	return srv
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer2(c *config.Service_Server) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Network != "" {
		opts = append(opts, http.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, http.Address(c.Addr2))
	}
	if c.Timeout != durationpb.New(0) {
		opts = append(opts, http.Timeout(c.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	return srv
}
