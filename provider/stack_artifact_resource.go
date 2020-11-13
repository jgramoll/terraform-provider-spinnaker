package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func artifactAccountResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"artifact_account": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"id": {
				Type:        schema.TypeString,
				Description: "",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"reference": {
				Type:        schema.TypeString,
				Description: "URL of file",
				Optional:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of artifact (github/file)",
				Required:    true,
			},
			"version": {
				Type:        schema.TypeString,
				Description: "Branch for github",
				Optional:    true,
			},
		},
	}
}
