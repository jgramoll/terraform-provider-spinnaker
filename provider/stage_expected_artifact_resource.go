package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func expectedArtifactResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"default_artifact": {
				Type:        schema.TypeList,
				Description: "Default Artifacts",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_kind": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
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
					},
				},
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
