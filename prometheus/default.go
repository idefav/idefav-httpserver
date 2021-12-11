package prometheus

import "github.com/prometheus/client_golang/prometheus"

var RequestRt *prometheus.HistogramVec

func init() {
	RequestRt = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "request_duration_seconds",
			Help: "the cost of request",
		},
		[]string{"Method", "Path", "Status"},
	)
	prometheus.MustRegister(RequestRt)
}
