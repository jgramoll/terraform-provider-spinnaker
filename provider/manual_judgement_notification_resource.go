package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func manualJudementNotificationResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
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
						"manual_judgment_continue": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"manual_judgment_stop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of notification (slack, email, etc)",
				Required:    true,
			},
			"when": &schema.Schema{
				Type:        schema.TypeList,
				Description: "When to send notification (started, completed, failed)",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manual_judgment": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"manual_judgment_continue": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"manual_judgment_stop": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}
