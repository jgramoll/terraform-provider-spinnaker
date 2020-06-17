package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineBakeManifestStageResource() *schema.Resource {
	newBakeManifestStageInterface := func() stage {
		return newBakeManifestStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newBakeManifestStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newBakeManifestStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newBakeManifestStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newBakeManifestStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"evaluate_override_expressions": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "",
				Optional:    true,
			},
			"input_artifact": &schema.Schema{
				Type:        schema.TypeList,
				Description: "",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"namespace": &schema.Schema{
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"output_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"overrides": &schema.Schema{
				Type:        schema.TypeMap,
				Description: "",
				Optional:    true,
			},
			"raw_overrides": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "",
				Optional:    true,
			},
			"template_renderer": &schema.Schema{
				Type:        schema.TypeString,
				Description: "",
				Required:    true,
			},
		}),
	}
}
