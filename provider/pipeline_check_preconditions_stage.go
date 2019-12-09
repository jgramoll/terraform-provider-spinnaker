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
	err := s.baseToClientStage(&cs.BaseStage, refID)
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

func (*checkPreconditionsStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.CheckPreconditionsStage)
	newStage := newCheckPreconditionsStage()
	newStage.baseFromClientStage(&clientStage.BaseStage)

	newStage.Preconditions = *fromClientPreconditions(&clientStage.Preconditions)

	return newStage
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
