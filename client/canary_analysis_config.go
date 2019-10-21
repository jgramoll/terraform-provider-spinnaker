package client

type CanaryAnalysisConfigScoreThreadholds struct {
	Marginal string `json:"marginal"`
	Pass     string `json:"pass"`
}

type CanaryAnalysisConfig struct {
	CanaryAnalysisIntervalMins string `json:"canaryAnalysisIntervalMins"`
	CanaryConfigId             string `json:"canaryConfigId"`
	LifetimeDuration           string `json:"lifetimeDuration"`
	MetricsAccountName         string `json:"metricsAccountName"`

	Scopes          []*CanaryAnalysisConfigScope         `json:"scopes"`
	ScoreThresholds CanaryAnalysisConfigScoreThreadholds `json:"scoreThresholds"`

	StorageAccountName string `json:"storageAccountName"`
}

func NewCanaryAnalysisConfig() *CanaryAnalysisConfig {
	return &CanaryAnalysisConfig{
		ScoreThresholds: CanaryAnalysisConfigScoreThreadholds{},
	}
}
