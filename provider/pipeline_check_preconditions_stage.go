package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type checkPreconditionsStage struct {
	baseStage `mapstructure:",squash"`

	Preconditions []precondition `mapstructure:"precondition"`
}

func newCheckPreconditionsStage() *checkPreconditionsStage {
	return &checkPreconditionsStage{
		baseStage: *newBaseStage(),
	}
}

func (s *checkPreconditionsStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewCheckPreconditionsStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	preconditions, err := toClientPreconditions(&s.Preconditions)
	if err != nil {
		return nil, err
	}
	cs.Preconditions = *preconditions

	return cs, nil
}

func (*checkPreconditionsStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.CheckPreconditionsStage)
	newStage := newCheckPreconditionsStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	preconditions, err := fromClientPreconditions(&clientStage.Preconditions)
	if err != nil {
		return nil, err
	}
	newStage.Preconditions = *preconditions

	return newStage, nil
}

func (s *checkPreconditionsStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("precondition", s.Preconditions)
	if err != nil {
		return err
	}

	return nil
}
