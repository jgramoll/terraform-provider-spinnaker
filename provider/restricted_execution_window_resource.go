package provider

import "github.com/hashicorp/terraform/helper/schema"

func restrictedExecutionWindowResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"days": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"jitter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"max_delay": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min_delay": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"skip_manual": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"whitelist": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_hour": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"end_min": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_hour": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_min": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}
