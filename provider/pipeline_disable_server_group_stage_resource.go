package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineDisableServerGroupStageResource() *schema.Resource {
	stageInterface := func() stage {
		return newDisableServerGroupStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, stageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, stageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, stageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, stageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},
		Schema: targetServerGroupStageResource(),
	}
}
