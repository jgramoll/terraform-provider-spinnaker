package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineCanaryAnalysisDeploymentsResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"baseline": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": &schema.Schema{
							Type:     schema.TypeMap,
							Required: true,
						},
						"application": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"cloud_provider": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"cluster": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"server_group_pair": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"control": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem:     clusterResource(),
						},
						"experiment": &schema.Schema{
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
