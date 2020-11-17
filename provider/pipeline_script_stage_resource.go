package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineScriptStageResource() *schema.Resource {
	newScriptStageInterface := func() stage {
		return newScriptStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newScriptStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newScriptStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newScriptStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newScriptStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"account": {
				Type:        schema.TypeString,
				Description: "The account name",
				Optional:    true,
			},
			"cluster": {
				Type:        schema.TypeString,
				Description: "The cluster to scale",
				Optional:    true,
			},
			"clusters": {
				Type:        schema.TypeList,
				Description: "The clusters to scale",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"command": {
				Type:        schema.TypeString,
				Description: "The criteria for determining the target cluster",
				Required:    true,
			},
			"cmc": {
				Type:        schema.TypeString,
				Description: "cmc passed down to script execution as CMC",
				Optional:    true,
			},
			"image": {
				Type:        schema.TypeString,
				Description: "The label selectors",
				Optional:    true,
			},
			"property_file": {
				Type:        schema.TypeString,
				Description: "The name to the properties file produced by the script execution to be used by later stages of the Spinnaker pipeline.",
				Optional:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Description: "The kind of manifest to scale",
				Optional:    true,
			},
			"regions": {
				Type:        schema.TypeList,
				Description: "The kinds of manifest to scale",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"repo_url": {
				Type:        schema.TypeString,
				Description: "Path to the repo hosting the scripts in Stash. (e.g. CDL/mimir-scripts). Leave empty to use the default.",
				Optional:    true,
			},
			"repo_branch": {
				Type:        schema.TypeString,
				Description: "Git Branch. (e.g. master). Leave empty to use the master branch.",
				Optional:    true,
			},
			"script_path": {
				Type:        schema.TypeString,
				Description: "Path to script to run",
				Required:    true,
			},
			"wait_for_completion": {
				Type:        schema.TypeBool,
				Description: "if false, marks the stage as successful right away without waiting for the script to complete",
				Optional:    true,
			},
		}),
	}
}
