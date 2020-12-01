package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type disableServerGroupStage struct {
	baseStage              `mapstructure:",squash"`
	targetServerGroupStage `mapstructure:",squash"`
}

func newDisableServerGroupStage() *disableServerGroupStage {
	return &disableServerGroupStage{
		baseStage: *newBaseStage(),
	}
}

func (s *disableServerGroupStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewDisableServerGroupStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}
	err = s.targetServerGroupStageToClient(&cs.TargetServerGroupStage)
	if err != nil {
		return nil, err
	}

	return cs, nil
}

func (*disableServerGroupStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.DisableServerGroupStage)
	newStage := newDisableServerGroupStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}
	err = newStage.targetServerGroupStageFromClientStage(&clientStage.TargetServerGroupStage)
	if err != nil {
		return nil, err
	}

	return newStage, nil
}

func (s *disableServerGroupStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}
	err = s.targetServerGroupSetResourceData(d)
	if err != nil {
		return err
	}

	return nil
}
