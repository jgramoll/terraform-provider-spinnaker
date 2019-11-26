package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineResizeServerGroupStageResource() *schema.Resource {
	stageInterface := func() stage {
		return newResizeServerGroupStage()
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
			"action": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Configures the resize action for the target server group (scale_down, scale_up, scale_to_cluster, scale_exact)",
				Required:    true,
			},
			"capacity": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Capacity for cluster",
				MaxItems:    1,
				Optional:    true,
				Elem:        capacityResource(),
			},
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
				Description: "Name of the cluster",
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
				Description: "regions to target (us-east-1)",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"resize_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of resize based on action (incremental, percentage)",
				Optional:    true,
			},
			"target": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Which server group to resize (ancestor_asg_dynamic, current_asg_dynamic, oldest_asg_dynamic)",
				Optional:    true,
			},
			"target_healthy_deploy_percentage": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Consider deploy successful when percent of instances are healthy",
				Optional:    true,
				Default:     100,
			},
		}),
	}
}
