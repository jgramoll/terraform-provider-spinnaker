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
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"ami_name": {
				Type:        schema.TypeString,
				Description: "Name of the ami output. Default = $package-$arch-$ami_suffix-$store_type",
				Optional:    true,
			},
			"ami_suffix": {
				Type:        schema.TypeString,
				Description: "Suffix of the ami output. String of date in format YYYYMMDDHHmm, default is calculated from timestamp",
				Optional:    true,
			},
			"base_ami": {
				Type:        schema.TypeString,
				Description: "Label of the base ami to use. If Base AMI is specified, this will be used instead of the Base OS provided",
				Optional:    true,
			},
			"base_label": {
				Type:        schema.TypeString,
				Description: "Label for the base ami (release)",
				Optional:    true,
				Default:     "release",
			},
			"base_name": {
				Type:        schema.TypeString,
				Description: "Name of the base ami to use",
				Optional:    true,
			},
			"base_os": {
				Type:        schema.TypeString,
				Description: "Base OS to use (trusty)",
				Optional:    true,
				Default:     "trusty",
			},
			"cloud_provider_type": {
				Type:        schema.TypeString,
				Description: "Cloud provider to use (aws)",
				Optional:    true,
			},
			"extended_attributes": {
				Type:        schema.TypeMap,
				Description: "Extra attributes to give the packer template",
				Optional:    true,
			},
			"package": {
				Type:        schema.TypeString,
				Description: "The name of the package you want installed (without any version identifiers).\nIf your build produces a deb file named \"myapp_1.27-h343\", you would want to enter \"myapp\" here.\nIf there are multiple packages (space separated), then they will be installed in the order they are entered.",
				Optional:    true,
			},
			"rebake": {
				Type:        schema.TypeBool,
				Description: "Rebake image without regard to the status of any existing bake",
				Optional:    true,
				Default:     false,
			},
			"regions": {
				Type:        schema.TypeList,
				Description: "regions to create the ami (us-east-1)",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"store_type": {
				Type:        schema.TypeString,
				Description: "Data store type to use when creating instances (ebs)",
				Optional:    true,
				Default:     "ebs",
			},
			"template_file_name": {
				Type:        schema.TypeString,
				Description: "Name of custom template to use",
				Optional:    true,
			},
			"var_file_name": {
				Type:        schema.TypeString,
				Description: "[Bakery] The name of a json file containing key/value pairs to add to the packer command",
				Optional:    true,
			},
			"vm_type": {
				Type:        schema.TypeString,
				Description: "Type of VM to use (hvm, pv)",
				Optional:    true,
				Default:     "hvm",
			},
		}),
	}
}
