package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type canaryAnalysisConfigs []*canaryAnalysisConfig

type canaryAnalysisConfig struct {
	CanaryAnalysisIntervalMins string `mapstructure:"canary_analysis_interval_mins"`
	CanaryConfigID             string `mapstructure:"canary_config_id"`
	LifetimeDuration           string `mapstructure:"lifetime_duration"`
	MetricsAccountName         string `mapstructure:"metrics_account_name"`

	Scopes          canaryAnalysisConfigScopes                `mapstructure:"scope"`
	ScoreThresholds canaryAnalysisConfigScoreThreadholdsArray `mapstructure:"score_thresholds"`

	StorageAccountName string `mapstructure:"storage_account_name"`
}

func (configs *canaryAnalysisConfigs) toClientCanaryConfig() *client.CanaryAnalysisConfig {
	for _, c := range *configs {
		return &client.CanaryAnalysisConfig{
			CanaryAnalysisIntervalMins: c.CanaryAnalysisIntervalMins,
			CanaryConfigID:             c.CanaryConfigID,
			LifetimeDuration:           c.LifetimeDuration,
			MetricsAccountName:         c.MetricsAccountName,
			Scopes:                     *c.Scopes.toClientCanaryConfigScopes(),
			ScoreThresholds:            *c.ScoreThresholds.toClientCanaryConfigScoreThresholds(),
			StorageAccountName:         c.StorageAccountName,
		}
	}
	return nil
}

func (*canaryAnalysisConfigs) fromClientCanaryConfig(c *client.CanaryAnalysisConfig) *canaryAnalysisConfigs {
	newConfig := canaryAnalysisConfig{
		CanaryAnalysisIntervalMins: c.CanaryAnalysisIntervalMins,
		CanaryConfigID:             c.CanaryConfigID,
		LifetimeDuration:           c.LifetimeDuration,
		MetricsAccountName:         c.MetricsAccountName,
		StorageAccountName:         c.StorageAccountName,
	}
	newConfig.Scopes = *newConfig.Scopes.fromClientCanaryConfigScopes(&c.Scopes)
	newConfig.ScoreThresholds = *newConfig.ScoreThresholds.fromClientCanaryConfigScoreThresholds(&c.ScoreThresholds)

	return &canaryAnalysisConfigs{&newConfig}
}
