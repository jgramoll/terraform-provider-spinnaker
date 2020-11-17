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
			"account": {
				Type:        schema.TypeString,
				Description: "The account name",
				Required:    true,
			},
			"application": {
				Type:        schema.TypeString,
				Description: "The application name",
				Optional:    true,
			},
			"cloud_provider": {
				Type:        schema.TypeString,
				Description: "The cloud provider name",
				Required:    true,
			},
			"cluster": {
				Type:        schema.TypeString,
				Description: "The cluster to scale",
				Optional:    true,
			},
			"criteria": {
				Type:        schema.TypeString,
				Description: "The criteria for determining the target cluster",
				Optional:    true,
			},
			"kind": {
				Type:        schema.TypeString,
				Description: "The kind of manifest to scale",
				Optional:    true,
			},
			"kinds": {
				Type:        schema.TypeList,
				Description: "The kinds of manifest to scale",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"label_selectors": {
				Type:        schema.TypeMap,
				Description: "The label selectors",
				Optional:    true,
			},
			"location": {
				Type:        schema.TypeString,
				Description: "The namespace",
				Optional:    true,
			},
			"manifest_name": {
				Type:        schema.TypeString,
				Description: "The name of the manifest",
				Optional:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "The selector mode",
				Optional:    true,
			},
			"replicas": {
				Type:        schema.TypeString,
				Description: "The number of replicas",
				Optional:    true,
			},
		}),
	}
}
