package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineDeployManifestStageResource() *schema.Resource {
	newDeployManifestStageInterface := func() stage {
		return newDeployManifestStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newDeployManifestStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newDeployManifestStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newDeployManifestStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newDeployManifestStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"account": {
				Type:        schema.TypeString,
				Description: "The account name",
				Required:    true,
			},
			"credentials": {
				Type:        schema.TypeString,
				Description: "The credentials name",
				Optional:    true,
			},
			"namespace_override": {
				Type:        schema.TypeString,
				Description: "Namespace override",
				Optional:    true,
			},
			"cloud_provider": {
				Type:        schema.TypeString,
				Description: "The cloud provider name",
				Required:    true,
			},
			"manifest_artifact_account": {
				Type:        schema.TypeString,
				Description: "The manifest artifact account",
				Optional:    true,
			},
			"manifest_artifact_id": {
				Type:        schema.TypeString,
				Description: "The manifest artifact id",
				Optional:    true,
			},
			"manifests": {
				Type:        schema.TypeList,
				Description: "The manifests as yaml",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"moniker": {
				Type:        schema.TypeList,
				Description: "Name to attach to manifest",
				Required:    true,
				MaxItems:    1,
				Elem:        monikerResource(),
			},
			"relationships": {
				Type:        schema.TypeList,
				Description: "relationships",
				Required:    true,
				MaxItems:    1,
				Elem:        relationshipsResource(),
			},
			"skip_expression_evaluation": {
				Type:        schema.TypeBool,
				Description: "Skip Expression Evaluation",
				Optional:    true,
				Default:     false,
			},
			"source": {
				Type:        schema.TypeString,
				Description: "source",
				Required:    true,
			},
			"traffic_management": {
				Type:        schema.TypeList,
				Description: "The manifest artifact account",
				Required:    true,
				MaxItems:    1,
				Elem:        trafficManagementResource(),
			},
		}),
	}
}
