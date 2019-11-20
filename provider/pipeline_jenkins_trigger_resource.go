package provider

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
)

var errInvalidTriggerImportKey = errors.New("Invalid import key, must be pipelineID_triggerID")

func pipelineJenkinsTriggerResource() *schema.Resource {
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
		},
	}
}
