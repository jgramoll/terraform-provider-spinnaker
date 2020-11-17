package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelinePipelineResource() *schema.Resource {
	newStageInterface := func() stage {
		return newPipelineStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"application": {
				Type:        schema.TypeString,
				Description: "Application with target pipeline",
				Required:    true,
			},
			"target_pipeline": {
				Type:        schema.TypeString,
				Description: "Target pipeline",
				Required:    true,
			},
			"pipeline_parameters": {
				Type:        schema.TypeMap,
				Description: "Parameters to pass to pipeline",
				Optional:    true,
			},
			"wait_for_completion": {
				Type:        schema.TypeBool,
				Description: "if false, marks the stage as successful right away without waiting for the pipeline to complete",
				Optional:    true,
				Default:     true,
			},
		}),
	}
}
