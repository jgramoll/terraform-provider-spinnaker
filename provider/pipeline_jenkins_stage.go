package provider

import (
	"strconv"

	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type jenkinsStage struct {
	baseStage `mapstructure:",squash"`

	Job                      string                 `mapstructure:"job"`
	MarkUnstableAsSuccessful bool                   `mapstructure:"mark_unstable_as_successful"`
	Master                   string                 `mapstructure:"master"`
	Parameters               map[string]interface{} `mapstructure:"parameters"`
	PropertyFile             string                 `mapstructure:"property_file"`
}

func newJenkinsStage() *jenkinsStage {
	return &jenkinsStage{
		baseStage: *newBaseStage(),
	}
}

func (s *jenkinsStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewJenkinsStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.Job = s.Job
	cs.MarkUnstableAsSuccessful = s.MarkUnstableAsSuccessful
	cs.Master = s.Master
	cs.Parameters = s.Parameters
	cs.PropertyFile = s.PropertyFile

	// hack around terraform not supporting booleans
	for key, value := range cs.Parameters {
		if v, ok := value.(string); ok {
			if v == "true" {
				cs.Parameters[key] = true
			} else if v == "false" {
				cs.Parameters[key] = false
			}
		}
	}

	return cs, nil
}

func (*jenkinsStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.JenkinsStage)
	newStage := newJenkinsStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.Job = clientStage.Job
	newStage.MarkUnstableAsSuccessful = clientStage.MarkUnstableAsSuccessful
	newStage.Master = clientStage.Master
	newStage.Parameters = clientStage.Parameters
	newStage.PropertyFile = clientStage.PropertyFile

	// hack around terraform not supporting booleans
	for key, value := range newStage.Parameters {
		if v, ok := value.(bool); ok {
			newStage.Parameters[key] = strconv.FormatBool(v)
		}
	}

	return newStage, nil
}

func (s *jenkinsStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("job", s.Job)
	if err != nil {
		return err
	}
	err = d.Set("mark_unstable_as_successful", s.MarkUnstableAsSuccessful)
	if err != nil {
		return err
	}
	err = d.Set("master", s.Master)
	if err != nil {
		return err
	}
	err = d.Set("parameters", s.Parameters)
	if err != nil {
		return err
	}
	return d.Set("property_file", s.PropertyFile)
}
