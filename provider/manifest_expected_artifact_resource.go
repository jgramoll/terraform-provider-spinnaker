package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func manifestExpectedArtifactResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"default_artifact": {
				Type:        schema.TypeList,
				Description: "Default Artifacts",
				Optional:    true,
				MaxItems:    1,
				Elem:        manifestArtifactResource(),
			},
			"display_name": {
				Type:        schema.TypeString,
				Description: "Name to display",
				Required:    true,
			},
			"id": {
				Type:        schema.TypeString,
				Description: "ID of the artifact",
				Computed:    true,
			},
			"match_artifact": {
				Type:        schema.TypeList,
				Description: "Artifact to match",
				Required:    true,
				MaxItems:    1,
				Elem:        manifestArtifactResource(),
			},
			"use_default_artifact": {
				Type:        schema.TypeBool,
				Description: "Use default artifact if missing",
				Default:     false,
				Optional:    true,
			},
			"use_prior_artifact": {
				Type:        schema.TypeBool,
				Description: "Use prior artifact if missing",
				Default:     false,
				Optional:    true,
			},
		},
	}
}
