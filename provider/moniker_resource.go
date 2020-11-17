package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func monikerResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"detail": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sequence": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stack": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
