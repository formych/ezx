package ezx

import (
	"net/http"

	"github.com/fsm-xyz/ezx/config"
	"github.com/fsm-xyz/ezx/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func metircs() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(config.C.Metrics.Prometheus.Addr, nil); err != nil {
		log.Std.Fatal("init metrics failed", zap.Error(err))
	}
}
