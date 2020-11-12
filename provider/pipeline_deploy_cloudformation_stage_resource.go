package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineDeployCloudformationStageResource() *schema.Resource {
	newDeployCloudformationStageInterface := func() stage {
		return newDeployCloudformationStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newDeployCloudformationStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newDeployCloudformationStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newDeployCloudformationStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newDeployCloudformationStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"action_on_replacement": {
				Type:        schema.TypeString,
				Description: "Action to take if ChangeSet contains a replacement",
				Default:     "ask",
				Optional:    true,
			},
			"capabilities": {
				Type:        schema.TypeList,
				Description: "",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"change_set_name": {
				Type:        schema.TypeString,
				Description: "Name of ChangeSet",
				Optional:    true,
			},
			"credentials": {
				Type:        schema.TypeString,
				Description: "Name of AWS account to use",
				Required:    true,
			},
			"execute_change_set": {
				Type:        schema.TypeBool,
				Description: "If ChangeSet should be executed",
				Default:     false,
				Optional:    true,
			},
			"is_change_set": {
				Type:        schema.TypeBool,
				Description: "If ChangeSet should be created",
				Default:     false,
				Optional:    true,
			},
			"parameters": {
				Type:        schema.TypeMap,
				Description: "",
				Optional:    true,
			},
			"regions": {
				Type:        schema.TypeList,
				Description: "",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"role_arn": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"source": {
				Type: schema.TypeString,
				Description: "Where the template file content is read from." +

					"text: The template is supplied statically to the pipeline from the below text-box." +

					"artifact: The template is read from an artifact supplied/created upstream. " +
					"The expected artifact must be referenced here, and will be bound at runtime.",
				Default:  "text",
				Optional: true,
			},
			"stack_artifact": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Template to use if source is artifact",
				Optional:    true,
				Elem:        artifactAccountResource(),
			},
			"stack_name": {
				Type:        schema.TypeString,
				Description: "Stack Name",
				Required:    true,
			},
			"tags": {
				Type:        schema.TypeMap,
				Description: "",
				Optional:    true,
			},
			"template_body": {
				Type:        schema.TypeList,
				Description: "Template if source is text",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		}),
	}
}
