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
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				return resourcePipelineImporter(d, meta, newPipelineStage().SetResourceData)
			},
		},

		Schema: map[string]*schema.Schema{
			PipelineKey: &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the pipeline to send notification",
				Required:    true,
				ForceNew:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the stage",
				Required:    true,
			},
			"requisite_stage_ref_ids": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Stage(s) that must be complete before this one",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notification": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Notifications to send for stage results",
				Optional:    true,
				Elem:        notificationResource(),
			},
			"stage_enabled": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Stage will only execute when the supplied expression evaluates true.\nThe expression does not need to be wrapped in ${ and }.\nIf this expression evaluates to false, the stages following this stage will still execute.",
				Optional:    true,
				MaxItems:    1,
				Elem:        stageEnabledResource(),
			},
			"application": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Application with target pipeline",
				Required:    true,
			},
			"complete_other_branches_then_fail": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "halt this branch and fail the pipeline once other branches complete. Prevents any stages that depend on this stage from running, but allows other branches of the pipeline to run. The pipeline will be marked as failed once complete.",
				Optional:    true,
				Default:     false,
			},
			"continue_pipeline": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If false, marks the stage as successful right away without waiting for the pipeline to complete",
				Optional:    true,
				Default:     false,
			},
			"fail_pipeline": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If the stage fails, immediately halt execution of all running stages and fails the entire execution",
				Optional:    true,
				Default:     true,
			},
			"fail_on_failed_expressions": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "The stage will be marked as failed if it contains any failed expressions",
				Optional:    true,
				Default:     false,
			},
			"target_pipeline": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Target pipeline",
				Required:    true,
			},
			"pipeline_parameters": &schema.Schema{
				Type:        schema.TypeMap,
				Description: "Parameters to pass to pipeline",
				Optional:    true,
			},
			"wait_for_completion": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "if false, marks the stage as successful right away without waiting for the pipeline to complete",
				Optional:    true,
				Default:     true,
			},
		},
	}
}
