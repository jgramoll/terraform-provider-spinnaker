package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineTriggerResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineTriggerCreate,
		Read:   resourcePipelineTriggerRead,
		Update: resourcePipelineTriggerUpdate,
		Delete: resourcePipelineTriggerDelete,
		Importer: &schema.ResourceImporter{
			State: resourceTriggerImporter,
		},
		DeprecationMessage: "use spinnaker_pipeline_jenkins_trigger",

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
			"job": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the job",
				Required:    true,
			},
			"master": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the job master",
				Required:    true,
			},
			"property_file": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of file to use for properties",
				Optional:    true,
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of trigger (jenkins, etc)",
				Optional:    true,
			},
		},
	}
}
