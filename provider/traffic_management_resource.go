package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func trafficManagementResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"options": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_traffic": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"services": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"strategy": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}
