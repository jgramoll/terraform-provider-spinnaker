package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func relationshipsResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"load_balancers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"security_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
