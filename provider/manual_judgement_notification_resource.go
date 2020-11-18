package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func manualJudementNotificationResource() *schema.Resource {
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
						"manual_judgment_continue": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"manual_judgment_stop": {
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
						"manual_judgment": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"manual_judgment_continue": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"manual_judgment_stop": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}
