package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineCheckPreconditionsStageResource() *schema.Resource {
	newCheckPreconditionsStageInterface := func() stage {
		return newCheckPreconditionsStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newCheckPreconditionsStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newCheckPreconditionsStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newCheckPreconditionsStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newCheckPreconditionsStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"precondition": {
				Type:        schema.TypeList,
				Description: "The preconditions for the stage",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"context": {
							Type:        schema.TypeMap,
							Description: "Map to describe precondition",
							Required:    true,
						},
						"fail_pipeline": {
							Type:        schema.TypeBool,
							Description: "The pipeline will fail whenever this precondition is false",
							Optional:    true,
							Default:     true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "The type of precondition. (expression, stageStatus, etc)",
							Required:    true,
						},
					},
				},
			},
		}),
	}
}
