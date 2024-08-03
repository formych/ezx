package kratos

import (
	"os"

	"github.com/fsm-xyz/ezx/config"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"

	"go.opentelemetry.io/otel/trace"
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

var (
	httpServer *http.Server
	grpcServer *grpc.Server
)

func GetHTTPServer() *http.Server {
	return httpServer
}

func GetGRPCServer() *grpc.Server {
	return grpcServer
}

func Init() {
	if config.C.Server.Type == config.GRPCServerType {
		grpcServer = NewGRPCServer(config.C.Server)
		// 服务是grpc时默认开启http端口
		config.C.Server.Addr = config.C.Server.Addr2
	}
	// 默认当作http服务
	httpServer = NewHTTPServer(config.C.Server)
}

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

	zlog.Hook(LogHook{})
	var op kratos.Option
	if grpcServer != nil {
		op = kratos.Server(grpcServer)
		// 异步开启http监听
		go kratos.New(
			kratos.ID(id),
			kratos.Name(Name),
			kratos.Version(Version),
			kratos.Metadata(map[string]string{}),
			kratos.Logger(logger),
			kratos.Server(httpServer),
		).Run()
	} else {
		op = kratos.Server(httpServer)
	}

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		op,
	).Run()
}

type LogHook struct{}

func (LogHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	traceID := ""
	ctx := e.GetCtx()
	if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
		traceID = span.TraceID().String()
	}

	e.Str("trace", traceID)
}
