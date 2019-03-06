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
			"complete_other_branches_then_fail": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "halt this branch and fail the pipeline once other branches complete. Prevents any stages that depend on this stage from running, but allows other branches of the pipeline to run. The pipeline will be marked as failed once complete.",
				Optional:    true,
				Default:     false,
			},
			"continue_pipeline": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If false, marks the stage as successful right away without waiting for the jenkins job to complete",
				Optional:    true,
				Default:     false,
			},
			"fail_pipeline": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If the stage fails, immediately halt execution of all running stages and fails the entire execution",
				Optional:    true,
				Default:     true,
			},
			"fail_on_failed_expressions": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "The stage will be marked as failed if it contains any failed expressions",
				Optional:    true,
				Default:     false,
			},
			"override_timeout": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Allows you to override the amount of time the stage can run before failing.\nNote: this represents the overall time the stage has to complete (the sum of all the task times).",
				Optional:    true,
				Default:     false,
			},
			"restrict_execution_during_time_window": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Restrict execution to specific time windows",
				Optional:    true,
				Default:     false,
			},
			"restricted_execution_window": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Time windows to restrict execution",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"days": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"jitter": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
									},
									"max_delay": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"min_delay": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"skip_manual": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"whitelist": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_hour": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"end_min": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"start_hour": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"start_min": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
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
			"package": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the package you want installed (without any version identifiers).\nIf your build produces a deb file named \"myapp_1.27-h343\", you would want to enter \"myapp\" here.\nIf there are multiple packages (space separated), then they will be installed in the order they are entered.",
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
