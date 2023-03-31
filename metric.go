package ezx

import (
	"net/http"

	"github.com/fsm-xyz/ezx/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func metircs() error {
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(config.C.Metrics.Prometheus.Addr, nil)
}
