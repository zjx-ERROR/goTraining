package server

import (
	"github.com/google/wire"
	"github.com/prometheus/client_golang/prometheus"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer)

var _metricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "server",
	Subsystem: "requests",
	Name:      "duration_ms",
	Help:      "server requests duration(ms).",
	Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000},
}, []string{"kind", "operation"})

var _metricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace: "client",
	Subsystem: "requests",
	Name:      "code_total",
	Help:      "The total number of processed requests",
}, []string{"kind", "operation", "code", "reason"})
