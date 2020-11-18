package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineWebhookTriggerResource() *schema.Resource {
	triggerInterface := func() trigger {
		return newWebhookTrigger()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineTriggerCreate(d, m, triggerInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineTriggerRead(d, m, triggerInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineTriggerUpdate(d, m, triggerInterface)
		},
		Delete: resourcePipelineTriggerDelete,
		Importer: &schema.ResourceImporter{
			State: resourceTriggerImporter,
		},

		Schema: triggerResource(map[string]*schema.Schema{
			"source": {
				Type:        schema.TypeString,
				Description: "Name of the webhook source",
				Required:    true,
			},
			"payload_constraints": {
				Type:        schema.TypeMap,
				Description: "Payload contraints",
				Optional:    true,
			},
		}),
	}
}
