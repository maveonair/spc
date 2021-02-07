package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"
)

const (
	listeningAddr = "0.0.0.0:9100"
)

var (
	errors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "errors_total",
		Help: "Number of errors.",
	})

	releaseProbe = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "spc_probe_success",
		Help: "Displays whether or not the probe was a success.",
	}, []string{"name", "version"})
)

func init() {
	prometheus.MustRegister(errors)
	prometheus.MustRegister(releaseProbe)
}

// Setup setup metric endpoint
func Setup() {
	setupMetricsEndpoint()
}

func setupMetricsEndpoint() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(listeningAddr, nil))
	}()
}

// IncreaseErrors increase errors counter
func IncreaseErrors() {
	errors.Inc()
}

// SetReleaseSuccessProbe sets the probe status for a release
func SetReleaseSuccessProbe(name string, version string, value float64) {
	releaseProbe.WithLabelValues(name, version).Set(value)
}
