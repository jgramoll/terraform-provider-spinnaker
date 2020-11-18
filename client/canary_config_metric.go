package client

// CanaryConfigMetric config metric
type CanaryConfigMetric struct {
	// AnalysisConfigurations interface `json:"analysisConfigurations"`
	Name      string                   `json:"name"`
	Query     *CanaryConfigMetricQuery `json:"query"`
	Groups    []string                 `json:"groups"`
	ScopeName string                   `json:"scopeName"`
}

// NewCanaryConfigMetric new config metric
func NewCanaryConfigMetric(
	group string, name string, metricQuery *CanaryConfigMetricQuery,
) *CanaryConfigMetric {
	return &CanaryConfigMetric{
		Name:      name,
		Query:     metricQuery,
		Groups:    []string{group},
		ScopeName: "default",
	}
}
