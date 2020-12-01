package provider

import (
	"errors"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

var errInvalidTriggerImportKey = errors.New("Invalid import key, must be pipelineID_triggerID")

func triggerResource(in map[string]*schema.Schema) map[string]*schema.Schema {
	out := map[string]*schema.Schema{
		PipelineKey: {
			Type:        schema.TypeString,
			Description: "Id of the pipeline to trigger",
			Required:    true,
			ForceNew:    true,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "If the trigger is enabled",
			Optional:    true,
			Default:     true,
		},
	}

	// merge input
	for k, v := range in {
		out[k] = v
	}

	return out
}

func resourcePipelineTriggerCreate(d *schema.ResourceData, m interface{}, createTrigger func() trigger) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	t := createTrigger()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, t); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	clientTrigger, err := t.toClientTrigger(id.String())
	if err != nil {
		return err
	}
	pipeline.AppendTrigger(clientTrigger)

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating pipeline trigger: %s\n", id)
	d.SetId(id.String())
	return resourcePipelineTriggerRead(d, m, createTrigger)
}

func resourcePipelineTriggerRead(d *schema.ResourceData, m interface{}, createTrigger func() trigger) error {
	pipelineID := d.Get(PipelineKey).(string)
	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(pipelineID)
	if err != nil {
		log.Printf("[WARN] No Pipeline found: %s\n", err)
		d.SetId("")
		return nil
	}

	clientTrigger, err := pipeline.GetTrigger(d.Id())
	if err != nil {
		log.Printf("[WARN] No Pipeline Trigger found: %s\n", err)
		d.SetId("")
		return nil
	}

	t, err := createTrigger().fromClientTrigger(clientTrigger)
	if err != nil {
		log.Printf("[WARN] No Pipeline Trigger found: %s\n", err)
		d.SetId("")
		return nil
	}
	return t.setResourceData(d)
}

func resourcePipelineTriggerUpdate(d *schema.ResourceData, m interface{}, createTrigger func() trigger) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	t := createTrigger()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, t); err != nil {
		return err
	}

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	clientTrigger, err := t.toClientTrigger(d.Id())
	if err != nil {
		return err
	}

	err = pipeline.UpdateTrigger(clientTrigger)
	if err != nil {
		return err
	}

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updated pipeline trigger: %s\n", d.Id())
	return resourcePipelineTriggerRead(d, m, createTrigger)
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
	log.Printf("[INFO] Importing d: %s\n", d.Get(""))
	log.Printf("[INFO] Importing id: %s\n", d.Id())
	id := strings.Split(d.Id(), "_")
	if len(id) < 2 {
		return nil, errInvalidTriggerImportKey
	}
	d.Set(PipelineKey, id[0])
	d.SetId(id[1])
	return []*schema.ResourceData{d}, nil
}
