package ezx

import (
	"net/http"

	"github.com/fsm-xyz/ezx/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

func metircs() {
	if config.C.Metrics == nil {
		return
	}
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(config.C.Metrics.Prometheus.Addr, nil); err != nil {
		log.Fatal().Err(err).Msg("init metrics failed")
	}
}
