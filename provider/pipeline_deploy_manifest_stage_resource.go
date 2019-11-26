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
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The account name",
				Required:    true,
			},
			"namespace_override": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Namespace override",
				Optional:    true,
			},
			"cloud_provider": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The cloud provider name",
				Required:    true,
			},
			"manifest_artifact_account": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The manifest artifact account",
				Optional:    true,
				Default:     "docker-registry",
			},
			"manifests": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The manifests as yaml",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"moniker": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Name to attach to manifest",
				Optional:    true,
				MaxItems:    1,
				Elem:        monikerResource(),
			},
			"relationships": &schema.Schema{
				Type:        schema.TypeList,
				Description: "relationships",
				Required:    true,
				MaxItems:    1,
				Elem:        relationshipsResource(),
			},
			"skip_expression_evaluation": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Skip Expression Evaluation",
				Optional:    true,
				Default:     false,
			},
			"source": &schema.Schema{
				Type:        schema.TypeString,
				Description: "source",
				Required:    true,
			},
			"traffic_management": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The manifest artifact account",
				Required:    true,
				MaxItems:    1,
				Elem:        trafficManagementResource(),
			},
		}),
	}
}
