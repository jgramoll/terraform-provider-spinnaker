package provider

// import (
// 	"github.com/hashicorp/terraform/helper/schema"
// )

// func pipelineJenkinsStageResource() *schema.Resource {
// 	stageType := "jenkins"
// 	return &schema.Resource{
// 		Create: func(d *schema.ResourceData, m interface{}) error {
// 			return resourcePipelineStageCreate(d, m, stageType)
// 		},
// 		Read: func(d *schema.ResourceData, m interface{}) error {
// 			return resourcePipelineStageRead(d, m, stageType)
// 		},
// 		Update: func(d *schema.ResourceData, m interface{}) error {
// 			return resourcePipelineStageUpdate(d, m, stageType)
// 		},
// 		Delete: func(d *schema.ResourceData, m interface{}) error {
// 			return resourcePipelineStageDelete(d, m, stageType)
// 		},

// 		Schema: map[string]*schema.Schema{
// 			"pipeline": &schema.Schema{
// 				Type:        schema.TypeString,
// 				Description: "Id of the pipeline to send notification",
// 				Required:    true,
// 				ForceNew:    true,
// 			},
// 			"completeOtherBranchesThenFail": &schema.Schema{
// 				Type:        schema.TypeBool,
// 				Description: "halt this branch and fail the pipeline once other branches complete. Prevents any stages that depend on this stage from running, but allows other branches of the pipeline to run. The pipeline will be marked as failed once complete.",
// 				Optional:    true,
// 				Default:     false,
// 			},
// 			"continuePipeline": &schema.Schema{
// 				Type:        schema.TypeBool,
// 				Description: "If false, marks the stage as successful right away without waiting for the jenkins job to complete",
// 				Optional:    true,
// 				Default:     false,
// 			},
// 			"failPipeline": &schema.Schema{
// 				Type:        schema.TypeBool,
// 				Description: "If the stage fails, immediately halt execution of all running stages and fails the entire execution",
// 				Optional:    true,
// 				Default:     true,
// 			},
// 			"job": &schema.Schema{
// 				Type:        schema.TypeString,
// 				Description: "Name of the Jenkins job to execute",
// 				Required:    true,
// 			},
// 			"markUnstableAsSuccessful": &schema.Schema{
// 				Type:        schema.TypeString,
// 				Description: "If Jenkins reports the build status as UNSTABLE, Spinnaker will mark the stage as SUCCEEDED and continue execution of the pipeline",
// 				Optional:    true,
// 				Default:     false,
// 			},
// 			"master": &schema.Schema{
// 				Type:        schema.TypeString,
// 				Description: "Name of the Jenkins master where the job will be executed",
// 				Required:    true,
// 			},
// 			"name": &schema.Schema{
// 				Type:        schema.TypeString,
// 				Description: "Name of the stage",
// 				Required:    true,
// 			},
// 			"parameters": &schema.Schema{
// 				Type:        schema.TypeMap,
// 				Description: "Parameters to pass to the Jenkins job",
// 				Optional:    true,
// 			},
// 			"propertyFile": &schema.Schema{
// 				Type:        schema.TypeString,
// 				Description: "Name of the property file to use for results",
// 				Optional:    true,
// 			},
// 		},
// 	}
// }
