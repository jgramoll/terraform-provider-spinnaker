package provider

import (
	"errors"
	"log"
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
			ApplicationKey: &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the application where the pipeline lives",
				Required:    true,
			},
			"disabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If the pipeline is disabled",
				Optional:    true,
				Default:     false,
			},
			"keep_waiting_pipelines": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Do not automatically cancel pipelines waiting in queue",
				Optional:    true,
				Default:     false,
			},
			"limit_concurrent": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Disable concurrent pipeline executions (only run one at a time)",
				Optional:    true,
				Default:     true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the pipeline",
				Required:    true,
			},
			"index": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Index of the pipeline",
				Optional:    true,
				Default:     0,
			},
			"roles": &schema.Schema{
				Type:        schema.TypeList,
				Description: "When the pipeline is triggered using an automated trigger, these roles will be used to decide if the pipeline has permissions to access a protected application or account.\n\nTo read from a protected application or account, the pipeline must have at least one role that has read access to the application or account.\nTo write to a protected application or account, the pipeline must have at least one role that has write access to the application or account.\nNote: To prevent privilege escalation vulnerabilities, a user must be a member of all of the groups specified here in order to modify, and execute the pipeline.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service_account": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Service account to run pipeline",
				Optional:    true,
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

	log.Println("[DEBUG] Creating pipeline:", pipeline.Application, pipeline.Name)
	pipelineService := m.(*Services).PipelineService
	err := pipelineService.CreatePipeline(&pipeline)
	if err != nil {
		return err
	}

	pipelineWithID, err := pipelineService.GetPipeline(pipeline.Application, pipeline.Name)
	if err != nil {
		log.Println("[WARN] No Pipeline found:", err)
		return err
	}

	log.Println("[DEBUG] New pipeline ID", pipelineWithID.ID)
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

	log.Printf("[INFO] Got Pipeline %s", p.ID)
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

	log.Println("[DEBUG] Updated pipeline:", d.Id())
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

	log.Println("[DEBUG] Deleting pipeline:", d.Id())
	pipelineService := m.(*Services).PipelineService
	return pipelineService.DeletePipeline(p.toClientPipeline())
}
