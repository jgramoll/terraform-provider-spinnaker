package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineEvaluateVariablesStageResource() *schema.Resource {
	newEvaluateVariablesStageInterface := func() stage {
		return newEvaluateVariablesStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newEvaluateVariablesStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newEvaluateVariablesStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newEvaluateVariablesStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newEvaluateVariablesStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"variables": {
				Type:        schema.TypeMap,
				Description: "List of values to assign as variables",
				Required:    true,
			},
		}),
	}
}
