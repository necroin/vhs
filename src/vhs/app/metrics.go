package app

import "github.com/necroin/golibs/libs/metrics"

type Metrics struct {
	Registry *metrics.Registry
	Storage  *MetricsStorage
}

type MetricsStorage struct {
	CallTimeHistogramVector *metrics.HistogramVector
}

func NewMetricsStorage() *MetricsStorage {
	return &MetricsStorage{
		CallTimeHistogramVector: metrics.NewHistogramVector(
			metrics.HistogramOpts{
				Name:    "call_time_histogram",
				Help:    "plugin service call time in ms",
				Buckets: metrics.Buckets{Start: 0, Range: 20, Count: 10},
			},
			"plugin", "service",
		),
	}
}

func NewMetrics() *Metrics {
	registry := metrics.NewRegistry()
	storage := NewMetricsStorage()

	registry.Register(storage.CallTimeHistogramVector)

	return &Metrics{
		Registry: registry,
		Storage:  storage,
	}
}
