package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineFindImageFromTagsStageResource() *schema.Resource {
	newFindImageFromTagsStageInterface := func() stage {
		return newFindImageFromTagsStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newFindImageFromTagsStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newFindImageFromTagsStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newFindImageFromTagsStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newFindImageFromTagsStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"cloud_provider": {
				Type:        schema.TypeString,
				Description: "Cloud provider to use (aws)",
				Required:    true,
			},
			"cloud_provider_type": {
				Type:        schema.TypeString,
				Description: "Cloud provider to use (aws)",
				Required:    true,
			},
			"package_name": {
				Type:        schema.TypeString,
				Description: "Label of the base ami to use. If Base AMI is specified, this will be used instead of the Base OS provided",
				Optional:    true,
			},
			"regions": {
				Type:        schema.TypeList,
				Description: "regions to target (us-east-1)",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": {
				Type:        schema.TypeMap,
				Description: "Tags of base ami to use.",
				Optional:    true,
			},
		}),
	}
}
