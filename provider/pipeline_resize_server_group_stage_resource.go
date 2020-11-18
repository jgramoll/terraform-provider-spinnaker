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
			"action": {
				Type:        schema.TypeString,
				Description: "Configures the resize action for the target server group (scale_down, scale_up, scale_to_cluster, scale_exact)",
				Required:    true,
			},
			"capacity": {
				Type:        schema.TypeList,
				Description: "Capacity for cluster",
				MaxItems:    1,
				Optional:    true,
				Elem:        capacityResource(),
			},
			"cloud_provider": {
				Type:        schema.TypeString,
				Description: "Cloud provider to use (aws)",
				Optional:    true,
			},
			"cloud_provider_type": {
				Type:        schema.TypeString,
				Description: "Cloud provider to use (aws)",
				Optional:    true,
			},
			"cluster": {
				Type:        schema.TypeString,
				Description: "Name of the cluster",
				Required:    true,
			},
			"credentials": {
				Type:        schema.TypeString,
				Description: "Name of the credentials to use",
				Optional:    true,
			},
			"moniker": {
				Type:        schema.TypeList,
				Description: "Name to attach to cluster",
				Optional:    true,
				MaxItems:    1,
				Elem:        monikerResource(),
			},
			"regions": {
				Type:        schema.TypeList,
				Description: "regions to target (us-east-1)",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"resize_type": {
				Type:        schema.TypeString,
				Description: "Type of resize based on action (incremental, percentage)",
				Optional:    true,
			},
			"target": {
				Type:        schema.TypeString,
				Description: "Which server group to resize (ancestor_asg_dynamic, current_asg_dynamic, oldest_asg_dynamic)",
				Optional:    true,
			},
			"target_healthy_deploy_percentage": {
				Type:        schema.TypeInt,
				Description: "Consider deploy successful when percent of instances are healthy",
				Optional:    true,
				Default:     100,
			},
		}),
	}
}
