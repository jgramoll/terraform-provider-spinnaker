package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineEnableServerGroupStageResource() *schema.Resource {
	stageInterface := func() stage {
		return newEnableServerGroupStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, stageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, stageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, stageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, stageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"cloud_provider": {
				Type:        schema.TypeString,
				Description: "Cloud provider to use (aws, kubernetes)",
				Required:    true,
			},
			"cloud_provider_type": {
				Type:        schema.TypeString,
				Description: "Cloud provider to use (aws, kubernetes)",
				Required:    true,
			},
			"cluster": {
				Type:        schema.TypeString,
				Description: "Cluster to enable",
				Required:    true,
			},
			"credentials": {
				Type:        schema.TypeString,
				Description: "Credentials to use with cloud provider",
				Required:    true,
			},
			"interesting_health_provider_names": {
				Type:        schema.TypeList,
				Description: "Health provider names",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"namespaces": {
				Type:        schema.TypeList,
				Description: "Deploy to K8s Namespaces",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"regions": {
				Type:        schema.TypeList,
				Description: "Deploy to AWS regions",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target": {
				Type:        schema.TypeString,
				Description: "Which version of cluster to target",
				Required:    true,
			},
		}),
	}
}
