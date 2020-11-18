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
			"evaluate_override_expressions": {
				Type:        schema.TypeBool,
				Description: "",
				Optional:    true,
			},
			"input_artifact": {
				Type:        schema.TypeList,
				Description: "",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"artifact": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem:     manifestArtifactResource(),
						},
					},
				},
			},
			"namespace": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"output_name": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"overrides": {
				Type:        schema.TypeMap,
				Description: "",
				Optional:    true,
			},
			"raw_overrides": {
				Type:        schema.TypeBool,
				Description: "",
				Optional:    true,
			},
			"template_renderer": {
				Type:        schema.TypeString,
				Description: "",
				Required:    true,
			},
			"kustomize_file_path": {
				Type:        schema.TypeString,
				Description: "Path to kustomization file (if using kustomize engine)",
				Optional:    true,
			},
		}),
	}
}
