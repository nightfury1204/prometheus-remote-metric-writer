package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	up = prometheus.NewDesc(
		prometheus.BuildFQName("", "", "up"),
		"Is metrics exporter is running.",
		nil, prometheus.Labels{
			"version": "dev",
			"who" : "metric-exporter",
		},
	)
)

type OperatorHealthCollector struct {
}

func NewOperatorHealthCollector() *OperatorHealthCollector {
	return &OperatorHealthCollector{}
}

func (c *OperatorHealthCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
}

func (c *OperatorHealthCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 1)
}
