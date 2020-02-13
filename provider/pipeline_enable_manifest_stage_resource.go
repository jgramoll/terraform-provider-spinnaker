package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineEnableManifestStageResource() *schema.Resource {
	newEnableManifestStageInterface := func() stage {
		return newEnableManifestStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newEnableManifestStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newEnableManifestStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newEnableManifestStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newEnableManifestStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The account name",
				Required:    true,
			},
			"app": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The application name",
				Required:    true,
			},
			"cloud_provider": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The cloud provider name",
				Required:    true,
			},
			"cluster": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the manifest to enable (e.g. replicatSet my-service)",
				Required:    true,
			},
			"criteria": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The target cluster (e.g. newest)",
				Required:    true,
			},
			"kind": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The cluster kind (e.g. replicaSet)",
				Required:    true,
			},
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The location name",
				Required:    true,
			},
			"mode": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The mode name",
				Required:    true,
			},
		}),
	}
}
