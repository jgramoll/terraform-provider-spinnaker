package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type canaryAnalysisStage struct {
	baseStage `mapstructure:",squash"`

	AnalysisType string                 `mapstructure:"analysis_type"`
	CanaryConfig *canaryAnalysisConfigs `mapstructure:"canary_config"`
	Deployments  *deployStageClusters   `mapstructure:"deployments"`
}

func newCanaryAnalysisStage() *canaryAnalysisStage {
	return &canaryAnalysisStage{
		baseStage:    *newBaseStage(),
		CanaryConfig: &canaryAnalysisConfigs{},
		Deployments:  &deployStageClusters{},
	}
}

func (s *canaryAnalysisStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewCanaryAnalysisStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.AnalysisType = s.AnalysisType
	cs.CanaryConfig = s.CanaryConfig.toClientCanaryConfig()
	if s.Deployments != nil {
		cs.Deployments = s.Deployments.toClientClusters()
	}

	return cs, nil
}

func (*canaryAnalysisStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.CanaryAnalysisStage)
	newStage := newCanaryAnalysisStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.AnalysisType = clientStage.AnalysisType
	newStage.CanaryConfig = newStage.CanaryConfig.fromClientCanaryConfig(clientStage.CanaryConfig)
	newStage.Deployments = newStage.Deployments.fromClientClusters(clientStage.Deployments)

	return newStage, nil
}

func (s *canaryAnalysisStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("analysis_type", s.AnalysisType)
	if err != nil {
		return err
	}
	err = d.Set("canary_config", s.CanaryConfig)
	if err != nil {
		return err
	}
	err = d.Set("deployments", s.Deployments)
	if err != nil {
		return err
	}
	return nil
}
