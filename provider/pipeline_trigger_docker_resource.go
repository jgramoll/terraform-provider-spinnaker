package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineDockerTriggerResource() *schema.Resource {
	triggerInterface := func() trigger {
		return newDockerTrigger()
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
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the account",
				Required:    true,
			},
			"organization": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the organization",
				Required:    true,
			},
			"registry": &schema.Schema{
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"repository": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of repository",
				Required:    true,
			},
			"tag": &schema.Schema{
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
		}),
	}
}
