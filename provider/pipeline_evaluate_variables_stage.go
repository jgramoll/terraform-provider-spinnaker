package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type evaluateVariablesStage struct {
	baseStage `mapstructure:",squash"`

	Variables map[string]string `mapstructure:"variables"`
}

func newEvaluateVariablesStage() *evaluateVariablesStage {
	return &evaluateVariablesStage{
		baseStage: *newBaseStage(),
	}
}

func (s *evaluateVariablesStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewEvaluateVariablesStage()
	err := s.baseToClientStage(&cs.BaseStage, refID)
	if err != nil {
		return nil, err
	}

	for k, v := range s.Variables {
		clientVariable := client.Variable{
			Key:   k,
			Value: v,
		}
		cs.Variables = append(cs.Variables, clientVariable)
	}

	return cs, nil
}

func (*evaluateVariablesStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.EvaluateVariablesStage)
	newStage := newEvaluateVariablesStage()
	newStage.baseFromClientStage(&clientStage.BaseStage)

	for _, clientVariable := range clientStage.Variables {
		newStage.Variables[clientVariable.Key] = clientVariable.Value
	}

	return newStage
}

func (s *evaluateVariablesStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("variables", s.Variables)
	if err != nil {
		return err
	}

	return nil
}
