package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type evaluateVariablesStage struct {
	baseStage `mapstructure:",squash"`

	Variables map[string]string `mapstructure:"variables"`
}

func newEvaluateVariablesStage() *evaluateVariablesStage {
	return &evaluateVariablesStage{
		baseStage: *newBaseStage(),
		Variables: make(map[string]string),
	}
}

func (s *evaluateVariablesStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewEvaluateVariablesStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
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

func (*evaluateVariablesStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.EvaluateVariablesStage)
	newStage := newEvaluateVariablesStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	for _, clientVariable := range clientStage.Variables {
		newStage.Variables[clientVariable.Key] = clientVariable.Value
	}

	return newStage, nil
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
