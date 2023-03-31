package config

import (
	"flag"
	"time"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"xyz-zyx.co/data/dbx"
	"xyz-zyx.co/data/rdbx"
	"xyz-zyx.co/log"
)

var (
	// C 全局配置
	C = &Service{}
	// flagconf is the config flag.
	flagconf string
)

const (
	HTTPServerType = "http"
	GRPCServerType = "grpc"
)

// Service 全局配置
type Service struct {
	Server  Server     `json:"server" yaml:"server"`
	Log     log.Config `json:"log" yaml:"log"`
	Metrics struct {
		Prometheus struct {
			Addr string `yaml:"addr"`
		} `yaml:"prometheus"`
	} `yaml:"metrics"`
	Data struct {
		DB    []dbx.Config  `json:"db" yaml:"db"`
		Redis []rdbx.Config `json:"redis" yaml:"redis"`
	} `yaml:"data"`
}

type Server struct {
	Name     int           `yaml:"name"`
	Type     string        `yaml:"type"`
	Network  string        `yaml:"network"`
	Addr     string        `yaml:"addr"`
	Timeout  time.Duration `yaml:"timeout"`
	Provider string        `yaml:"provider"`
}

// Kafka   Kafka            `json:"kafka" yaml:"kafka"`
// Clients []client.Config  `json:"clients" yaml:"clients"`
// Monitor Monitor          `json:"monitor" yaml:"monitor"`
// CronJob CronJob          `json:"cronjob" yaml:"cronjob"`

// CronJob 定时任务配置信息
// type CronJob struct {
// 	Spec string `json:"spec" yaml:"spec"`
// }

// Kafka ...
// type Kafka struct {
// 	Consumers []kafka.ReaderConfig `json:"consumers" yaml:"consumers"`
// 	Producers []kafka.WriterConfig `json:"producers" yaml:"producers"`
// }

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func Init(bc any) {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	if err := c.Scan(&C); err != nil {
		panic(err)
	}

	// 业务自定义数据的解析
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
}
