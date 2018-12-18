package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineBakeStageResource() *schema.Resource {
	stageType := "bake"
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newBakeStage)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			// TODO set factory
			return resourcePipelineStageRead(d, m, stageType)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			// TODO set factory
			return resourcePipelineStageUpdate(d, m, stageType)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			// TODO set factory
			return resourcePipelineStageDelete(d, m, stageType)
		},

		Schema: map[string]*schema.Schema{
			"pipeline": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the pipeline to send notification",
				Required:    true,
				ForceNew:    true,
			},
			"ami_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the ami output. Default = $package-$arch-$ami_suffix-$store_type",
				Optional:    true,
			},
			"ami_suffix": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Suffix of the ami output. String of date in format YYYYMMDDHHmm, default is calculated from timestamp",
				Optional:    true,
			},
			"base_ami": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Label of the base ami to use. If Base AMI is specified, this will be used instead of the Base OS provided",
				Optional:    true,
			},
			"base_label": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Label for the base ami (release)",
				Optional:    true,
			},
			"base_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the base ami to use",
				Optional:    true,
			},
			"base_os": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Base OS to use (trusty)",
				Optional:    true,
			},
			"cloud_provider_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Cloud provider to use (aws)",
				Optional:    true,
			},
			"extended_attributes": &schema.Schema{
				Type:        schema.TypeMap,
				Description: "Extra attributes to give the packer template",
				Optional:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the stage",
				Required:    true,
			},
			"regions": &schema.Schema{
				Type:        schema.TypeList,
				Description: "regions to create the ami (us-east-1)",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"requisite_stage_ref_ids": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Stage(s) that must be complete before this one",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"store_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Data store type to use when creating instances (ebs)",
				Optional:    true,
			},
			"template_file_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of custom template to use",
				Optional:    true,
			},
			"var_file_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "[Bakery] The name of a json file containing key/value pairs to add to the packer command",
				Optional:    true,
			},
			"vm_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of VM to use (hvm, pv)",
				Optional:    true,
			},
		},
	}
}
