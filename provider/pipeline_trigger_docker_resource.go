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

		Schema: map[string]*schema.Schema{
			PipelineKey: &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the pipeline to trigger",
				Required:    true,
				ForceNew:    true,
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If the trigger is enabled",
				Optional:    true,
				Default:     true,
			},
			"run_as_user": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of user to run pipeline as",
				Optional:    true,
			},
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
			"repository": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of repository",
				Required:    true,
			},
		},
	}
}
