package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineUndoRolloutManifestStageResource() *schema.Resource {
	stageInterface := func() stage {
		return newUndoRolloutManifestStage()
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

		Schema: stageResource(map[string]*schema.Schema{
			"account": {
				Type:        schema.TypeString,
				Description: "Spinnaker account to use",
				Required:    true,
			},
			"cloud_provider": {
				Type:        schema.TypeString,
				Description: "Cloud provider to use (aws, kubernetes)",
				Required:    true,
			},
			"location": {
				Type:        schema.TypeString,
				Description: "K8s namespace with manifest",
				Required:    true,
			},
			"manifest_name": {
				Type:        schema.TypeString,
				Description: "K8s manifest name with kind (e.g. replicaSet my-service)",
				Required:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "Rollback mode (static)",
				Optional:    true,
				Default:     "static",
			},
			"num_revisions_back": {
				Type:        schema.TypeInt,
				Description: "How many revisions to rollback",
				Required:    true,
			},
		}),
	}
}
