package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineManualJudgementStageResource() *schema.Resource {
	newManualJudgmentStageInterface := func() stage {
		return newManualJudgmentStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newManualJudgmentStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newManualJudgmentStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newManualJudgmentStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newManualJudgmentStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"instructions": {
				Type:        schema.TypeString,
				Description: "Instructions",
				Optional:    true,
			},
			"judgment_inputs": {
				Type:        schema.TypeList,
				Description: "Judgment Inputs",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notification": {
				Type:        schema.TypeList,
				Description: "Notifications to send for stage results",
				Optional:    true,
				Elem:        manualJudementNotificationResource(),
			},
		}),
	}
}
