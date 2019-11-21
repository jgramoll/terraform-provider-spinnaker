package provider

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

// Jenkins trigger for Pipeline
type jenkinsTrigger struct {
	Enabled   bool   `mapstructure:"enabled"`
	RunAsUser string `mapstructure:"run_as_user"`

	Job          string `mapstructure:"job"`
	Master       string `mapstructure:"master"`
	PropertyFile string `mapstructure:"property_file"`
}

func newJenkinsTrigger() *jenkinsTrigger {
	return &jenkinsTrigger{}
}

func (t *jenkinsTrigger) toClientTrigger(id string) (client.Trigger, error) {
	clientTrigger := client.NewJenkinsTrigger()
	clientTrigger.ID = id
	clientTrigger.Enabled = t.Enabled
	clientTrigger.Job = t.Job
	clientTrigger.Master = t.Master
	clientTrigger.PropertyFile = t.PropertyFile
	clientTrigger.RunAsUser = t.RunAsUser
	return clientTrigger, nil
}

func (*jenkinsTrigger) fromClientTrigger(clientTriggerInterface client.Trigger) (trigger, error) {
	clientTrigger, ok := clientTriggerInterface.(*client.JenkinsTrigger)
	if !ok {
		return nil, errors.New("Expected jenkins trigger")
	}
	t := newJenkinsTrigger()
	t.Enabled = clientTrigger.Enabled
	t.Job = clientTrigger.Job
	t.Master = clientTrigger.Master
	t.PropertyFile = clientTrigger.PropertyFile
	t.RunAsUser = clientTrigger.RunAsUser
	return t, nil
}

func (t *jenkinsTrigger) setResourceData(d *schema.ResourceData) error {
	var err error
	err = d.Set("enabled", t.Enabled)
	if err != nil {
		return err
	}
	err = d.Set("run_as_user", t.RunAsUser)
	if err != nil {
		return err
	}
	err = d.Set("job", t.Job)
	if err != nil {
		return err
	}
	err = d.Set("master", t.Master)
	if err != nil {
		return err
	}
	err = d.Set("property_file", t.PropertyFile)
	if err != nil {
		return err
	}
	return nil
}
