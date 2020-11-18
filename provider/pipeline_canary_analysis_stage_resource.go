package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineCanaryAnalysisStageResource() *schema.Resource {
	newCanaryAnalysisStageInterface := func() stage {
		return newCanaryAnalysisStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newCanaryAnalysisStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newCanaryAnalysisStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newCanaryAnalysisStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newCanaryAnalysisStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"analysis_type": {
				Type: schema.TypeString,
				Description: "Real Time analysis will be performed over a time interval beginning at the moment of execution." +
					"" +
					"Automatic: Spinnaker will provision and clean up the baseline and canary server groups. Not all cloud providers support this mode." +
					"Manual: You are responsible for provisioning and cleaning up the baseline and canary server groups." +
					"Retrospective analysis will be performed over an explicitly-specified time interval (likely in the past). You are responsible for provisioning and cleaning up the baseline and canary server groups.",
				Required: true,
			},
			"canary_config": {
				Type:        schema.TypeList,
				Description: "The manifest artifact account",
				Required:    true,
				MaxItems:    1,
				Elem:        pipelineCanaryAnalysisConfigResource(),
			},
			"deployments": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     pipelineCanaryAnalysisDeploymentsResource(),
			},
		}),
	}
}
