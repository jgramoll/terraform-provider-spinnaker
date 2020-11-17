package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineFindArtifactsFromResourceStageResource() *schema.Resource {
	newFindArtifactsFromResourceStageInterface := func() stage {
		return newFindArtifactsFromResourceStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newFindArtifactsFromResourceStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newFindArtifactsFromResourceStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newFindArtifactsFromResourceStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newFindArtifactsFromResourceStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"account": {
				Type:        schema.TypeString,
				Description: "Spinnaker account for cloud provider",
				Required:    true,
			},
			"cloud_provider": {
				Type:        schema.TypeString,
				Description: "Cloud provider to use (kubernetes)",
				Required:    true,
			},
			"location": {
				Type:        schema.TypeString,
				Description: "Location in cloud provider to search (k8s namespaces)",
				Required:    true,
			},
			"manifest_name": {
				Type:        schema.TypeString,
				Description: "If mode is static, resource manifest name",
				Optional:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "Manifest Selector (static, dynamic)",
				Required:    true,
			},
		}),
	}
}
