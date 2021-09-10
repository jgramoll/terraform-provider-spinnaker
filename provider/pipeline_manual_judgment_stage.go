package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type manualJudgmentStage struct {
	baseStage `mapstructure:",squash"`

	Instructions   string   `mapstructure:"instructions"`
	JudgmentInputs []string `mapstructure:"judgment_inputs"`
}

func newManualJudgmentStage() *manualJudgmentStage {
	return &manualJudgmentStage{
		baseStage: *newBaseStage(),
	}
}

func (s *manualJudgmentStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewManualJudgmentStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newManualJudgementNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.Instructions = s.Instructions

	judgmentInputs := []client.JudgmentInputs{}
	for _, v := range s.JudgmentInputs {
		input := client.JudgmentInputs{
			Value: v,
		}
		judgmentInputs = append(judgmentInputs, input)
	}
	cs.JudgmentInputs = judgmentInputs

	return cs, nil
}

func (*manualJudgmentStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.ManualJudgmentStage)
	newStage := newManualJudgmentStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newManualJudgementNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.Instructions = clientStage.Instructions

	var judgmentInputs []string
	for _, j := range clientStage.JudgmentInputs {
		judgmentInputs = append(judgmentInputs, j.Value)
	}
	newStage.JudgmentInputs = judgmentInputs

	return newStage, nil
}

func (s *manualJudgmentStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("instructions", s.Instructions)
	if err != nil {
		return err
	}
	return d.Set("judgment_inputs", s.JudgmentInputs)
}
