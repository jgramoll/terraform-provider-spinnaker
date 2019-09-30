package client

type CanaryConfigMetricQuery struct {
	Type        string `json:"type"`
	ServiceType string `json:"serviceType"`
	MetricName  string `json:"metricName"`
}

func NewCanaryConfigMetricQuery(metric string, serviceType string, metricType string) *CanaryConfigMetricQuery {
	return &CanaryConfigMetricQuery{
		Type:        metricType,
		ServiceType: serviceType,
		MetricName:  metric,
	}
}
