package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func capacityResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"desired": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"max": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"min": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
