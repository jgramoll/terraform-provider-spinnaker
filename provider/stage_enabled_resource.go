package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func stageEnabledResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"expression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
