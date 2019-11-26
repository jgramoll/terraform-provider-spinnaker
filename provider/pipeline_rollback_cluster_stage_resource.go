package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineRollbackClusterStageResource() *schema.Resource {
	newRollbackClusterInterface := func() stage {
		return newRollbackClusterStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newRollbackClusterInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newRollbackClusterInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newRollbackClusterInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newRollbackClusterInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"cloud_provider": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Cloud provider to use (aws)",
				Optional:    true,
			},
			"cloud_provider_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Cloud provider to use (aws)",
				Optional:    true,
			},
			"cluster": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the cluster to be rollback",
				Required:    true,
			},
			"credentials": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the credentials to use",
				Optional:    true,
			},
			"moniker": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Name to attach to cluster",
				Optional:    true,
				MaxItems:    1,
				Elem:        monikerResource(),
			},
			"regions": &schema.Schema{
				Type:        schema.TypeList,
				Description: "regions to rollback (us-east-1)",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_healthy_rollback_percentage": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Consider rollback successful when percent of instances are healthy",
				Optional:    true,
				Default:     100,
			},
		}),
	}
}
