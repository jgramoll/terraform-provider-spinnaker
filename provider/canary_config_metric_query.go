package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type canaryConfigMetricQueries []*canaryConfigMetricQuery

type canaryConfigMetricQuery struct {
	Type        string `mapstructure:"type"`
	ServiceType string `mapstructure:"service_type"`
	MetricName  string `mapstructure:"metric_name"`
}

func (q *canaryConfigMetricQueries) toClientMetricsQuery() *client.CanaryConfigMetricQuery {
	for _, query := range *q {
		return &client.CanaryConfigMetricQuery{
			Type:        query.Type,
			ServiceType: query.ServiceType,
			MetricName:  query.MetricName,
		}
	}
	return nil
}

func (*canaryConfigMetricQueries) fromClientMetricsQuery(c *client.CanaryConfigMetricQuery) *canaryConfigMetricQueries {
	var query canaryConfigMetricQuery
	query.Type = c.Type
	query.ServiceType = c.ServiceType
	query.MetricName = c.MetricName
	return &canaryConfigMetricQueries{&query}
}
