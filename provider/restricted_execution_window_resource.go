package provider

import "github.com/hashicorp/terraform/helper/schema"

func restrictedExecutionWindowResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"days": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"jitter": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"max_delay": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min_delay": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"skip_manual": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"whitelist": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_hour": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"end_min": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_hour": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_min": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}
