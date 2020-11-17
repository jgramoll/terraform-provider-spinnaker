package client

// CanaryConfigMetricQuery metric query
type CanaryConfigMetricQuery struct {
	Type        string `json:"type"`
	ServiceType string `json:"serviceType"`
	MetricName  string `json:"metricName"`
}

// NewCanaryConfigMetricQuery new metric query
func NewCanaryConfigMetricQuery(metric string, serviceType string, metricType string) *CanaryConfigMetricQuery {
	return &CanaryConfigMetricQuery{
		Type:        metricType,
		ServiceType: serviceType,
		MetricName:  metric,
	}
}
