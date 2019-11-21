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
			"triggering_application": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the spinnaker application",
				Required:    true,
			},
			"triggering_pipeline": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the spinnaker pipeline",
				Required:    true,
			},
			"status": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Status of pipeline execution",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
