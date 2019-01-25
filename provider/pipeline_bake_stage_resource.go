package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineBakeStageResource() *schema.Resource {
	newBakeStageInterface := func() stage {
		return newBakeStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newBakeStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newBakeStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newBakeStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newBakeStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			PipelineKey: &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the pipeline to send notification",
				Required:    true,
				ForceNew:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the stage",
				Required:    true,
			},
			"requisite_stage_ref_ids": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Stage(s) that must be complete before this one",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notification": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Notifications to send for stage results",
				Optional:    true,
				Elem:        notificationResource(),
			},
			"stage_enabled": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Stage will only execute when the supplied expression evaluates true.\nThe expression does not need to be wrapped in ${ and }.\nIf this expression evaluates to false, the stages following this stage will still execute.",
				Optional:    true,
				MaxItems:    1,
				Elem:        stageEnabledResource(),
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
				Default:     "release",
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
				Default:     "trusty",
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
			"rebake": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Rebake image without regard to the status of any existing bake",
				Optional:    true,
				Default:     false,
			},
			"regions": &schema.Schema{
				Type:        schema.TypeList,
				Description: "regions to create the ami (us-east-1)",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"store_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Data store type to use when creating instances (ebs)",
				Optional:    true,
				Default:     "ebs",
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
				Default:     "hvm",
			},
		},
	}
}
