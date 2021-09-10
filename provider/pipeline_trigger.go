package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type trigger interface {
	toClientTrigger(string) (client.Trigger, error)
	fromClientTrigger(client.Trigger) (trigger, error)
	setResourceData(d *schema.ResourceData) error
}

type baseTrigger struct {
	Enabled bool `mapstructure:"enabled"`
}
