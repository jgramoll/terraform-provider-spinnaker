package provider

import (
	"errors"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineDataSource() *schema.Resource {
	return &schema.Resource{
		Read: pipelineDataSourceRead,

		Schema: map[string]*schema.Schema{
			ApplicationKey: {
				Type:        schema.TypeString,
				Description: "Name of the application where the pipeline lives",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the pipeline",
				Required:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if !pipelineNameRegex.MatchString(v) {
						errs = append(errs, errors.New("Pipeline name cannot contain any of the following characters: / \\ ? % #"))
					}
					return
				},
			},
			"disabled": {
				Type:        schema.TypeBool,
				Description: "If the pipeline is disabled",
				Computed:    true,
			},
			"keep_waiting_pipelines": {
				Type:        schema.TypeBool,
				Description: "Do not automatically cancel pipelines waiting in queue",
				Computed:    true,
			},
			"limit_concurrent": {
				Type:        schema.TypeBool,
				Description: "Disable concurrent pipeline executions (only run one at a time)",
				Computed:    true,
			},
			"index": {
				Type:        schema.TypeInt,
				Description: "Index of the pipeline",
				Computed:    true,
			},
			"roles": {
				Type:        schema.TypeList,
				Description: "When the pipeline is triggered using an automated trigger, these roles will be used to decide if the pipeline has permissions to access a protected application or account.\n\nTo read from a protected application or account, the pipeline must have at least one role that has read access to the application or account.\nTo write to a protected application or account, the pipeline must have at least one role that has write access to the application or account.\nNote: To prevent privilege escalation vulnerabilities, a user must be a member of all of the groups specified here in order to modify, and execute the pipeline.",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service_account": {
				Type:        schema.TypeString,
				Description: "Service account to run pipeline",
				Computed:    true,
			},
			"locked": {
				Type:        schema.TypeList,
				Description: "Lock options",
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ui": {
							Type:        schema.TypeBool,
							Description: "Lock user to edit pipeline over the spinnaker UI",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Description banner explaining why ui is locked",
							Computed:    true,
						},
						"allow_unlock_ui": {
							Type:        schema.TypeBool,
							Description: "Allow user to unlock ui to edit pipeline",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func pipelineDataSourceRead(d *schema.ResourceData, m interface{}) error {
	application := d.Get(ApplicationKey).(string)
	name := d.Get("name").(string)

	log.Printf("[DEBUG] Importing pipeline %s on application %s\n", name, application)
	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipeline(application, name)
	if err != nil {
		log.Printf("[WARN] No Pipeline found: %v\n", err)
		return err
	}

	log.Printf("[DEBUG] Imported pipeline: %s\n", pipeline.ID)
	d.SetId(pipeline.ID)

	return resourcePipelineRead(d, m)
}
