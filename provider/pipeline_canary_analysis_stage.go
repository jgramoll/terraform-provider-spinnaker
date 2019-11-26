package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
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

func (s *canaryAnalysisStage) toClientStage(config *client.Config, refId string) (client.Stage, error) {
	cs := client.NewCanaryAnalysisStage()
	err := s.baseToClientStage(&cs.BaseStage, refId)
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

func (s *canaryAnalysisStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.CanaryAnalysisStage)
	newStage := newCanaryAnalysisStage()
	newStage.baseFromClientStage(&clientStage.BaseStage)

	newStage.AnalysisType = clientStage.AnalysisType
	newStage.CanaryConfig = newStage.CanaryConfig.fromClientCanaryConfig(clientStage.CanaryConfig)
	newStage.Deployments = newStage.Deployments.fromClientClusters(clientStage.Deployments)

	return newStage
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
