package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type pipelineTrigger interface {
	toClientTrigger(string) (client.Trigger, error)
	fromClientTrigger(client.Trigger) (pipelineTrigger, error)
	setResourceData(d *schema.ResourceData) error
}
