package provider

import "github.com/hashicorp/terraform/helper/schema"

func targetServerGroupStageResource() map[string]*schema.Schema {
	return stageResource(map[string]*schema.Schema{
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
		"target": {
			Type:        schema.TypeString,
			Description: "Which server group to destroy (oldest_asg_dynamic, ancestor_asg_dynamic, current_asg_dynamic)",
			Optional:    true,
		},
	})
}
