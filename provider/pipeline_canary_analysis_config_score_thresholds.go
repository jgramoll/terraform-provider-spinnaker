package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type canaryAnalysisConfigScoreThreadholdsArray []*canaryAnalysisConfigScoreThreadholds

type canaryAnalysisConfigScoreThreadholds struct {
	Marginal string `mapstructure:"marginal"`
	Pass     string `mapstructure:"pass"`
}

func (arr *canaryAnalysisConfigScoreThreadholdsArray) toClientCanaryConfigScoreThresholds() *client.CanaryAnalysisConfigScoreThreadholds {
	for _, c := range *arr {
		return &client.CanaryAnalysisConfigScoreThreadholds{
			Marginal: c.Marginal,
			Pass:     c.Pass,
		}
	}
	return nil
}

func (*canaryAnalysisConfigScoreThreadholdsArray) fromClientCanaryConfigScoreThresholds(
	c *client.CanaryAnalysisConfigScoreThreadholds,
) *canaryAnalysisConfigScoreThreadholdsArray {
	return &canaryAnalysisConfigScoreThreadholdsArray{
		&canaryAnalysisConfigScoreThreadholds{
			Marginal: c.Marginal,
			Pass:     c.Pass,
		},
	}
}
