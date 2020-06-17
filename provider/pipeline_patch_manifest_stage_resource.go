package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelinePatchManifestStageResource() *schema.Resource {
	newPatchManifestStageInterface := func() stage {
		return newPatchManifestStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newPatchManifestStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newPatchManifestStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newPatchManifestStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newPatchManifestStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"account": {
				Type:        schema.TypeString,
				Description: "",
				Required:    true,
			},
			"app": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"cloud_provider": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"cluster": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"criteria": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"kind": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"location": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"manifest_name": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"options": {
				Type:        schema.TypeList,
				Description: "",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"merge_strategy": {
							Type:     schema.TypeString,
							Required: true,
						},
						"record": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"patch_body": {
				Type:        schema.TypeList,
				Description: "",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
		}),
	}
}
