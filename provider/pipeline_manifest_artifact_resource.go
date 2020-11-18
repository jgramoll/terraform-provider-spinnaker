package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func manifestArtifactResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"artifact_account": {
				Type:        schema.TypeString,
				Description: "Account of artifact",
				Optional:    true,
			},
			"custom_kind": {
				Type:        schema.TypeBool,
				Description: "Artifact is custom kind",
				Optional:    true,
				Default:     false,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metadata": {
				Type:        schema.TypeMap,
				Description: "Metadata",
				Optional:    true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reference": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
