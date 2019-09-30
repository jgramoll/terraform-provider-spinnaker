package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func canaryConfigClassifierResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"group_weights": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}
