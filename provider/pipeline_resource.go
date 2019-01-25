package provider

import (
	"errors"
	"log"
	"sync"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

const ApplicationKey = "application"

var pipelineLock sync.Mutex

// ErrMissingPipelineName missing pipeline name
var ErrMissingPipelineName = errors.New("pipeline name must be provided")

// ErrMissingPipelineApplication missing pipeline application
var ErrMissingPipelineApplication = errors.New("pipeline application must be provided")

func pipelineResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineCreate,
		Read:   resourcePipelineRead,
		Update: resourcePipelineUpdate,
		Delete: resourcePipelineDelete,

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
			"parameter": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Pipeline parameters",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Default value",
							Optional:    true,
							Default:     false,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"label": &schema.Schema{
							Type:        schema.TypeString,
							Description: "A label to display when users are triggering the pipeline manually",
							Optional:    true,
							Default:     false,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"option": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"required": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
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
	pipeline, err := pipelineService.GetPipelineByID(d.Id())
	if err != nil {
		log.Println("[WARN] No Pipeline found:", d.Id())
		d.SetId("")
		return nil
	}

	log.Printf("[INFO] Got Pipeline %s", pipeline.ID)
	SetResourceData(pipeline, d)
	return nil
}

func resourcePipelineUpdate(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Id())
	if err != nil {
		return err
	}
	PipelineFromResourceData(pipeline, d)

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated pipeline:", d.Id())
	return resourcePipelineRead(d, m)
}

func resourcePipelineDelete(d *schema.ResourceData, m interface{}) error {
	var pipeline Pipeline
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &pipeline); err != nil {
		return err
	}

	if pipeline.Name == "" {
		return ErrMissingPipelineName
	}
	if pipeline.Application == "" {
		return ErrMissingPipelineApplication
	}

	log.Println("[DEBUG] Deleting pipeline:", d.Id())
	d.SetId("")
	pipelineService := m.(*Services).PipelineService
	return pipelineService.DeletePipeline(pipeline.ToClientPipeline())
}
