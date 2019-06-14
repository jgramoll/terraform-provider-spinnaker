package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func capacityResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"desired": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"desired_expression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_expression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_expression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
