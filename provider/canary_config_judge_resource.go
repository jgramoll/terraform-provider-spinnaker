package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func canaryConfigJudgeResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
