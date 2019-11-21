package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineJenkinsTriggerResource(deprecationMessage string) *schema.Resource {
	triggerInterface := func() trigger {
		return newJenkinsTrigger()
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
		DeprecationMessage: deprecationMessage,

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
				Description: "[DEPRECATED] Type of trigger, not used use explicit trigger resource",
				Optional:    true,
				Deprecated:  "DO NOT USE",
			},
		},
	}
}
