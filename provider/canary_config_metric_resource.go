package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func canaryConfigMetricResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"query": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem:     canaryConfigMetricQueryResource(),
			},
			"groups": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"scope_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
		},
	}
}
