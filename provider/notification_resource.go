package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func notificationResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Description: "Address of the notification (slack channel, email, etc)",
				Required:    true,
			},
			"message": {
				Type:        schema.TypeList,
				Description: "Custom messages",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"complete": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"failed": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"starting": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of notification (slack, email, etc)",
				Required:    true,
			},
			"when": {
				Type:        schema.TypeList,
				Description: "When to send notification (started, completed, failed)",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"complete": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"failed": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"starting": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}
