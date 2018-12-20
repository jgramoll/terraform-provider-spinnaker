package provider

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

// ErrTriggerNotFound trigger not found
var ErrTriggerNotFound = errors.New("could not find trigger")

// Trigger for Pipeline
type Trigger struct {
	ID           string
	Enabled      bool
	Job          string
	Master       string
	PropertyFile string `mapstructure:"property_file"`
	Type         string
}

func pipelineTriggerResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineTriggerCreate,
		Read:   resourcePipelineTriggerRead,
		Update: resourcePipelineTriggerUpdate,
		Delete: resourcePipelineTriggerDelete,

		Schema: map[string]*schema.Schema{
			PipelineKey: &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the pipeline to trigger",
				Required:    true,
				ForceNew:    true,
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If the trigger is enabled",
				Optional:    true,
				Default:     true,
			},
			"job": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the job",
				Required:    true,
			},
			"master": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the job master",
				Required:    true,
			},
			"property_file": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of file to use for properties",
				Optional:    true,
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of trigger (jenkins, etc)",
				Required:    true,
			},
		},
	}
}

func resourcePipelineTriggerCreate(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	var trigger Trigger
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &trigger); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	trigger.ID = id.String()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	pipeline.Triggers = append(pipeline.Triggers, client.Trigger(trigger))
	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Creating pipeline trigger:", id)
	d.SetId(id.String())
	return resourcePipelineTriggerRead(d, m)
}

func resourcePipelineTriggerRead(d *schema.ResourceData, m interface{}) error {
	pipelineID := d.Get(PipelineKey).(string)
	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(pipelineID)
	if err != nil {
		log.Println("[WARN] No Pipeline found:", err)
		d.SetId("")
		return nil
	}

	var trigger *client.Trigger
	trigger, err = getTrigger(pipeline.Triggers, d.Id())
	if err != nil {
		log.Println("[WARN] No Pipeline Trigger found:", err)
		d.SetId("")
	} else {
		d.SetId(trigger.ID)
		d.Set("enabled", trigger.Enabled)
		d.Set("job", trigger.Job)
		d.Set("master", trigger.Master)
		d.Set("property_file", trigger.PropertyFile)
		d.Set("type", trigger.Type)
	}

	return nil
}

func getTrigger(triggers []client.Trigger, triggerID string) (*client.Trigger, error) {
	for _, trigger := range triggers {
		if trigger.ID == triggerID {
			return &trigger, nil
		}
	}
	return nil, ErrTriggerNotFound
}

func resourcePipelineTriggerUpdate(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	var trigger Trigger
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &trigger); err != nil {
		return err
	}
	trigger.ID = d.Id()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	err = updateTriggers(pipeline, client.Trigger(trigger))
	if err != nil {
		return err
	}

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated pipeline trigger:", d.Id())
	return resourcePipelineTriggerRead(d, m)
}

func updateTriggers(pipeline *client.Pipeline, trigger client.Trigger) error {
	for i, t := range pipeline.Triggers {
		if t.ID == trigger.ID {
			pipeline.Triggers[i] = trigger
			return nil
		}
	}
	return ErrTriggerNotFound
}

func resourcePipelineTriggerDelete(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	var trigger Trigger
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &trigger); err != nil {
		return err
	}
	trigger.ID = d.Id()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	err = deleteTrigger(pipeline, &trigger)
	if err != nil {
		return err
	}

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func deleteTrigger(pipeline *client.Pipeline, trigger *Trigger) error {
	for i, t := range pipeline.Triggers {
		if t.ID == trigger.ID {
			pipeline.Triggers = append(pipeline.Triggers[:i], pipeline.Triggers[i+1:]...)
			return nil
		}
	}
	return ErrTriggerNotFound
}
