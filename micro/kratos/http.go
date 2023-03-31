package kratos

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"xyz-zyx.co/config"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *config.Server) *http.Server {
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
	if c.Timeout > 0 {
		opts = append(opts, http.Timeout(c.Timeout))
	}
	srv := http.NewServer(opts...)
	return srv
}
