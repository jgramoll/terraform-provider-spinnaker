package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

// Trigger for Pipeline
type trigger struct {
	ID           string
	Enabled      bool
	Job          string
	Master       string
	PropertyFile string `mapstructure:"property_file"`
	RunAsUser    string `mapstructure:"run_as_user"`
	Type         string
}

func fromClientTrigger(clientTrigger *client.Trigger) *trigger {
	t := trigger(*clientTrigger)
	return &t
}

func (t *trigger) setResourceData(d *schema.ResourceData) error {
	err := d.Set("enabled", t.Enabled)
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
	return d.Set("type", t.Type)
}
