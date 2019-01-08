package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func monikerResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"detail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"stack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
