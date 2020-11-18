package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineCanaryAnalysisDeploymentsResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"baseline": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": {
							Type:     schema.TypeMap,
							Required: true,
						},
						"application": {
							Type:     schema.TypeString,
							Required: true,
						},
						"cloud_provider": {
							Type:     schema.TypeString,
							Required: true,
						},
						"cluster": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"server_group_pair": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"control": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem:     clusterResource(),
						},
						"experiment": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem:     clusterResource(),
						},
					},
				},
			},
		},
	}
}
