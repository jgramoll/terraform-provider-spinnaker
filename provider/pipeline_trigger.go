package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type trigger interface {
	toClientTrigger(string) (client.Trigger, error)
	fromClientTrigger(client.Trigger) (trigger, error)
	setResourceData(d *schema.ResourceData) error
}

type baseTrigger struct {
	Enabled   bool   `mapstructure:"enabled"`
	RunAsUser string `mapstructure:"run_as_user"`
}
