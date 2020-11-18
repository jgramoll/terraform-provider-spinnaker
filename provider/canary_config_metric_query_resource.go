package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func canaryConfigMetricQueryResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metric_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
