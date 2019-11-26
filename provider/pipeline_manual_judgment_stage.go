package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
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

func (s *manualJudgmentStage) toClientStage(config *client.Config, refId string) (client.Stage, error) {
	cs := client.NewManualJudgmentStage()
	err := s.baseToClientStage(&cs.BaseStage, refId)
	if err != nil {
		return nil, err
	}

	cs.Instructions = s.Instructions

	var judgmentInputs []client.JudgmentInputs
	for _, v := range s.JudgmentInputs {
		judgmentInputs = append(judgmentInputs, client.JudgmentInputs{v})
	}
	cs.JudgmentInputs = judgmentInputs

	return cs, nil
}

func (s *manualJudgmentStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.ManualJudgmentStage)
	newStage := newManualJudgmentStage()
	newStage.baseFromClientStage(&clientStage.BaseStage)

	newStage.Instructions = clientStage.Instructions

	var judgmentInputs []string
	for _, j := range clientStage.JudgmentInputs {
		judgmentInputs = append(judgmentInputs, j.Value)
	}
	newStage.JudgmentInputs = judgmentInputs

	return newStage
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
