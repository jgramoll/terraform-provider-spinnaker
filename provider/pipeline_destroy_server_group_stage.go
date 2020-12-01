package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type destroyServerGroupStage struct {
	baseStage              `mapstructure:",squash"`
	targetServerGroupStage `mapstructure:",squash"`
}

func newDestroyServerGroupStage() *destroyServerGroupStage {
	return &destroyServerGroupStage{
		baseStage: *newBaseStage(),
	}
}

func (s *destroyServerGroupStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewDestroyServerGroupStage()
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

func (*destroyServerGroupStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.DestroyServerGroupStage)
	newStage := newDestroyServerGroupStage()
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

func (s *destroyServerGroupStage) SetResourceData(d *schema.ResourceData) error {
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
