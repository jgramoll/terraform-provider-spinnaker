package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type canaryConfigMetrics []*canaryConfigMetric

type canaryConfigMetric struct {
	// AnalysisConfigurations interface `mapstructure:"analysis_configurations"`
	Name      string                    `mapstructure:"name"`
	Query     canaryConfigMetricQueries `mapstructure:"query"`
	Groups    []string                  `mapstructure:"groups"`
	ScopeName string                    `mapstructure:"scope_name"`
}

func (m *canaryConfigMetrics) toClientMetrics() *[]*client.CanaryConfigMetric {
	metrics := []*client.CanaryConfigMetric{}
	for _, metric := range *m {
		metrics = append(metrics, &client.CanaryConfigMetric{
			Name:      metric.Name,
			Query:     metric.Query.toClientMetricsQuery(),
			Groups:    metric.Groups,
			ScopeName: metric.ScopeName,
		})
	}
	return &metrics
}

func (*canaryConfigMetrics) fromClientMetrics(c *[]*client.CanaryConfigMetric) *canaryConfigMetrics {
	metrics := canaryConfigMetrics{}
	for _, metric := range *c {
		newMetric := canaryConfigMetric{
			Name:      metric.Name,
			Groups:    metric.Groups,
			ScopeName: metric.ScopeName,
		}
		newMetric.Query = *newMetric.Query.fromClientMetricsQuery(metric.Query)
		metrics = append(metrics, &newMetric)
	}
	return &metrics
}
