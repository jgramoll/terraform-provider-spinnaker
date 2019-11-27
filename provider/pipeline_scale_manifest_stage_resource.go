package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineScaleManifestStageResource() *schema.Resource {
	newScaleManifestStageInterface := func() stage {
		return newScaleManifestStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newScaleManifestStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newScaleManifestStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newScaleManifestStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newScaleManifestStageInterface)
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
			"application": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The application name",
				Optional:    true,
			},
			"cloud_provider": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The cloud provider name",
				Required:    true,
			},
			"cluster": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The cluster to scale",
				Optional:    true,
			},
			"criteria": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The criteria for determining the target cluster",
				Optional:    true,
			},
			"kind": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The kind of manifest to scale",
				Optional:    true,
			},
			"kinds": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The kinds of manifest to scale",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"label_selectors": &schema.Schema{
				Type:        schema.TypeMap,
				Description: "The label selectors",
				Optional:    true,
			},
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The namespace",
				Optional:    true,
			},
			"manifest_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the manifest",
				Optional:    true,
			},
			"mode": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The selector mode",
				Optional:    true,
			},
			"replicas": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The number of replicas",
				Optional:    true,
			},
		}),
	}
}
