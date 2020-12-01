package client

// CanaryAnalysisConfigScoreThreadholds thresholds
type CanaryAnalysisConfigScoreThreadholds struct {
	Marginal string `json:"marginal"`
	Pass     string `json:"pass"`
}

// CanaryAnalysisConfig config
type CanaryAnalysisConfig struct {
	CanaryAnalysisIntervalMins string `json:"canaryAnalysisIntervalMins"`
	CanaryConfigID             string `json:"canaryConfigId"`
	LifetimeDuration           string `json:"lifetimeDuration"`
	MetricsAccountName         string `json:"metricsAccountName"`

	Scopes          []*CanaryAnalysisConfigScope         `json:"scopes"`
	ScoreThresholds CanaryAnalysisConfigScoreThreadholds `json:"scoreThresholds"`

	StorageAccountName string `json:"storageAccountName"`
}

// NewCanaryAnalysisConfig new config
func NewCanaryAnalysisConfig() *CanaryAnalysisConfig {
	return &CanaryAnalysisConfig{
		ScoreThresholds: CanaryAnalysisConfigScoreThreadholds{},
	}
}
