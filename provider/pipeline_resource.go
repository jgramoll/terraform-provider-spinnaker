package provider

import (
	"errors"
	"log"
	"regexp"
	"sync"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

const (
	// PipelineKey key for pipeline in map
	PipelineKey = "pipeline"
)

var (
	pipelineLock sync.Mutex

	pipelineNameRegex = regexp.MustCompile("^[a-zA-Z_0-9.][^\\?/%#]*$")

	// ErrMissingPipelineName missing pipeline name
	ErrMissingPipelineName = errors.New("pipeline name must be provided")

	// ErrMissingPipelineApplication missing pipeline application
	ErrMissingPipelineApplication = errors.New("pipeline application must be provided")
)

func pipelineResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineCreate,
		Read:   resourcePipelineRead,
		Update: resourcePipelineUpdate,
		Delete: resourcePipelineDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			ApplicationKey: {
				Type:        schema.TypeString,
				Description: "Name of the application where the pipeline lives",
				Required:    true,
			},
			"disabled": {
				Type:        schema.TypeBool,
				Description: "If the pipeline is disabled",
				Optional:    true,
				Default:     false,
			},
			"keep_waiting_pipelines": {
				Type:        schema.TypeBool,
				Description: "Do not automatically cancel pipelines waiting in queue",
				Optional:    true,
				Default:     false,
			},
			"limit_concurrent": {
				Type:        schema.TypeBool,
				Description: "Disable concurrent pipeline executions (only run one at a time)",
				Optional:    true,
				Default:     true,
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
			"index": {
				Type:        schema.TypeInt,
				Description: "Index of the pipeline",
				Optional:    true,
				Computed:    true,
			},
			"roles": {
				Type:        schema.TypeList,
				Description: "When the pipeline is triggered using an automated trigger, these roles will be used to decide if the pipeline has permissions to access a protected application or account.\n\nTo read from a protected application or account, the pipeline must have at least one role that has read access to the application or account.\nTo write to a protected application or account, the pipeline must have at least one role that has write access to the application or account.\nNote: To prevent privilege escalation vulnerabilities, a user must be a member of all of the groups specified here in order to modify, and execute the pipeline.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service_account": {
				Type:        schema.TypeString,
				Description: "Service account to run pipeline",
				Optional:    true,
			},
			"locked": {
				Type:        schema.TypeList,
				Description: "Lock options",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ui": {
							Type:        schema.TypeBool,
							Description: "Lock user to edit pipeline over the spinnaker UI",
							Optional:    true,
							Default:     false,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Description banner explaining why ui is locked",
							Optional:    true,
						},
						"allow_unlock_ui": {
							Type:        schema.TypeBool,
							Description: "Allow user to unlock ui to edit pipeline",
							Optional:    true,
							Default:     true,
						},
					},
				},
			},
		},
	}
}

func resourcePipelineCreate(d *schema.ResourceData, m interface{}) error {
	var pipeline client.CreatePipelineRequest
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &pipeline); err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating pipeline %s on application %s\n", pipeline.Name, pipeline.Application)
	pipelineService := m.(*Services).PipelineService
	err := pipelineService.CreatePipeline(&pipeline)
	if err != nil {
		return err
	}

	pipelineWithID, err := pipelineService.GetPipeline(pipeline.Application, pipeline.Name)
	if err != nil {
		log.Printf("[WARN] No Pipeline found: %s\n", err)
		return err
	}

	log.Printf("[DEBUG] New pipeline ID %s\n", pipelineWithID.ID)
	d.SetId(pipelineWithID.ID)
	// create can't update index...
	return resourcePipelineUpdate(d, m)
}

func resourcePipelineRead(d *schema.ResourceData, m interface{}) error {
	pipelineService := m.(*Services).PipelineService
	p, err := pipelineService.GetPipelineByID(d.Id())
	if err != nil {
		if serr, ok := err.(*client.SpinnakerError); ok {
			if serr.Status == 404 {
				d.SetId("")
				return nil
			}
		}

		return err
	}

	log.Printf("[INFO] Got Pipeline: %s\n", p.ID)
	return fromClientPipeline(p).setResourceData(d)
}

func resourcePipelineUpdate(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Id())
	if err != nil {
		if serr, ok := err.(*client.SpinnakerError); ok {
			if serr.Status == 404 {
				d.SetId("")
				return nil
			}
		}

		return err
	}

	pipelineFromResourceData(pipeline, d)

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updated pipeline: %s\n", d.Id())
	return resourcePipelineRead(d, m)
}

func resourcePipelineDelete(d *schema.ResourceData, m interface{}) error {
	var p pipeline
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &p); err != nil {
		return err
	}

	if p.Name == "" {
		return ErrMissingPipelineName
	}
	if p.Application == "" {
		return ErrMissingPipelineApplication
	}

	log.Printf("[DEBUG] Deleting pipeline: %s\n", d.Id())
	pipelineService := m.(*Services).PipelineService
	return pipelineService.DeletePipeline(p.toClientPipeline())
}
