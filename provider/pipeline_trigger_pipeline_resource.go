package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelinePipelineTriggerResource() *schema.Resource {
	triggerInterface := func() trigger {
		return newPipelineTrigger()
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
			"triggering_application": {
				Type:        schema.TypeString,
				Description: "Name of the spinnaker application",
				Required:    true,
			},
			"triggering_pipeline": {
				Type:        schema.TypeString,
				Description: "Name of the spinnaker pipeline",
				Required:    true,
			},
			"status": {
				Type:        schema.TypeList,
				Description: "Status of pipeline execution",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		}),
	}
}
