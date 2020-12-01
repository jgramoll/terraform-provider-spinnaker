package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type canaryConfig struct {
	Name          string              `mapstructure:"name"`
	Applications  []string            `mapstructure:"applications"`
	Description   string              `mapstructure:"description"`
	Metrics       canaryConfigMetrics `mapstructure:"metric"`
	ConfigVersion string              `mapstructure:"config_version"`
	// Templates     map[string]interface{}    `mapstructure:"templates"`
	Classifier canaryConfigClassifiers `mapstructure:"classifier"`
	Judge      canaryConfigJudges      `mapstructure:"judge"`
}

func (c *canaryConfig) toClientCanaryConfig(id string) *client.CanaryConfig {
	return &client.CanaryConfig{
		ID:            id,
		Name:          c.Name,
		Applications:  c.Applications,
		Description:   c.Description,
		Metrics:       *c.Metrics.toClientMetrics(),
		ConfigVersion: c.ConfigVersion,
		// Templates:     c.Templates,
		Classifier: c.Classifier.toClientClassifier(),
		Judge:      c.Judge.toClientJudge(),
	}
}

func fromClientCanaryConfig(c *client.CanaryConfig) *canaryConfig {
	config := &canaryConfig{
		Name:          c.Name,
		Applications:  c.Applications,
		Description:   c.Description,
		ConfigVersion: c.ConfigVersion,
	}
	config.Metrics = *config.Metrics.fromClientMetrics(&c.Metrics)
	config.Classifier = *config.Classifier.fromClientClassifier(c.Classifier)
	config.Judge = *config.Judge.fromClientJudge(c.Judge)
	return config
}

func (c *canaryConfig) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("name", c.Name); err != nil {
		return err
	}
	if err := d.Set("applications", c.Applications); err != nil {
		return err
	}
	if err := d.Set("description", c.Description); err != nil {
		return err
	}
	if err := d.Set("metric", c.Metrics); err != nil {
		return err
	}
	if err := d.Set("config_version", c.ConfigVersion); err != nil {
		return err
	}
	// if err := d.Set("templates", c.Templates); err != nil {
	// 	return err
	// }
	if err := d.Set("classifier", c.Classifier); err != nil {
		return err
	}
	return d.Set("judge", c.Judge)
}
