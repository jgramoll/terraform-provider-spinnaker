package provider

import (
	"errors"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

var errInvalidTriggerImportKey = errors.New("Invalid import key, must be pipelineID_triggerID")

func pipelineTriggerResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineTriggerCreate,
		Read:   resourcePipelineTriggerRead,
		Update: resourcePipelineTriggerUpdate,
		Delete: resourcePipelineTriggerDelete,
		Importer: &schema.ResourceImporter{
			State: resourceTriggerImporter,
		},

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

	var t trigger
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &t); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	t.ID = id.String()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	clientTrigger := client.Trigger(t)
	pipeline.AppendTrigger(&clientTrigger)

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

	var clientTrigger *client.Trigger
	clientTrigger, err = pipeline.GetTrigger(d.Id())
	if err != nil {
		log.Println("[WARN] No Pipeline Trigger found:", err)
		d.SetId("")
	} else {
		d.SetId(clientTrigger.ID)
		fromClientTrigger(clientTrigger).setResourceData(d)
	}

	return nil
}

func resourcePipelineTriggerUpdate(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	var t trigger
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &t); err != nil {
		return err
	}
	t.ID = d.Id()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	clientTrigger := client.Trigger(t)
	err = pipeline.UpdateTrigger(&clientTrigger)
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

func resourcePipelineTriggerDelete(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	err = pipeline.DeleteTrigger(d.Id())
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

func resourceTriggerImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	log.Println("[INFO] Importing d: ", d.Get(""))
	log.Println("[INFO] Importing id: ", d.Id())
	id := strings.Split(d.Id(), "_")
	if len(id) < 2 {
		return nil, errInvalidTriggerImportKey
	}
	d.Set(PipelineKey, id[0])
	d.SetId(id[1])
	return []*schema.ResourceData{d}, nil
}
