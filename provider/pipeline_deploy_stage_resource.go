package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineDeployStageResource() *schema.Resource {
	newDeployStageInterface := func() stage {
		return newDeployStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newDeployStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newDeployStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newDeployStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newDeployStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"cluster": {
				Type:        schema.TypeList,
				Description: "The clusters to be deployed",
				Required:    true,
				Elem:        pipelineDeployStageClusterResource(),
			},
		}),
	}
}
