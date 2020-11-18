package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineDeleteManifestStageResource() *schema.Resource {
	newDeleteManifestStageInterface := func() stage {
		return newDeleteManifestStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newDeleteManifestStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newDeleteManifestStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newDeleteManifestStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newDeleteManifestStageInterface)
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
			"app": {
				Type:        schema.TypeString,
				Description: "The application name",
				Required:    true,
			},
			"cloud_provider": {
				Type:        schema.TypeString,
				Description: "The cloud provider name",
				Required:    true,
			},
			"location": {
				Type:        schema.TypeString,
				Description: "The location name",
				Required:    true,
			},
			"manifest_name": {
				Type:        schema.TypeString,
				Description: "The manifest name",
				Required:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "The mode name",
				Required:    true,
			},
			"options": {
				Type:        schema.TypeList,
				Description: "Options for delete",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cascading": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
		}),
	}
}
