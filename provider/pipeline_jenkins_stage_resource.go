package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineJenkinsStageResource() *schema.Resource {
	newJenkinsStageInterface := func() stage {
		return newJenkinsStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newJenkinsStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newJenkinsStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newJenkinsStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newJenkinsStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"job": {
				Type:        schema.TypeString,
				Description: "Name of the Jenkins job to execute",
				Required:    true,
			},
			"mark_unstable_as_successful": {
				Type:        schema.TypeBool,
				Description: "If Jenkins reports the build status as UNSTABLE, Spinnaker will mark the stage as SUCCEEDED and continue execution of the pipeline",
				Optional:    true,
				Default:     false,
			},
			"master": {
				Type:        schema.TypeString,
				Description: "Name of the Jenkins master where the job will be executed",
				Required:    true,
			},
			"parameters": {
				Type:        schema.TypeMap,
				Description: "Parameters to pass to the Jenkins job",
				Optional:    true,
			},
			"property_file": {
				Type:        schema.TypeString,
				Description: "Name of the property file to use for results",
				Optional:    true,
			},
		}),
	}
}
