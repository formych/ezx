package kratos

import (
	"os"

	"github.com/formych/ezx/config"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string

	id, _ = os.Hostname()
)

// cleanup()
func Run() error {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	var srv transport.Server
	if config.C.Server.Type == config.GRPCServerType {
		srv = NewGRPCServer(&config.C.Server)

	} else if config.C.Server.Type == config.HTTPServerType {
		srv = NewHTTPServer(&config.C.Server)
	}

	return newApp(logger, srv).Run()
}

func newApp(logger log.Logger, srv ...transport.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			srv...,
		),
	)
}
