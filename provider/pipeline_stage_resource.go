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

var errInvalidStageImportKey = errors.New("Invalid import key, must be pipelineID_stageID")

func resourcePipelineImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	id := strings.Split(d.Id(), "_")
	if len(id) < 2 {
		return nil, errInvalidStageImportKey
	}
	d.SetId(id[1])
	err := d.Set(PipelineKey, id[0])
	if err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func resourcePipelineStageCreate(d *schema.ResourceData, m interface{}, createStage func() stage) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	s := createStage()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &s); err != nil {
		return err
	}
	stage := s.(stage)

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	stage.SetRefID(id.String())

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	cs, err := stage.toClientStage(&m.(*Services).Config)
	if err != nil {
		return err
	}
	pipeline.AppendStage(cs)

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Creating pipeline stage:", id)
	d.SetId(id.String())
	return resourcePipelineStageRead(d, m, createStage)
}

func resourcePipelineStageRead(d *schema.ResourceData, m interface{}, createStage func() stage) error {
	pipelineID := d.Get(PipelineKey).(string)
	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(pipelineID)
	if err != nil {
		log.Println("[WARN] No Pipeline found:", err)
		d.SetId("")
		return nil
	}

	var cStage client.Stage
	cStage, err = pipeline.GetStage(d.Id())
	if err != nil {
		log.Println("[WARN] No Pipeline Stage found:", err)
		d.SetId("")
		return err
	}

	s := createStage().(stage)
	s = s.fromClientStage(cStage)
	log.Println("[INFO] Updating from stage", cStage)
	return s.SetResourceData(d)

}

func resourcePipelineStageUpdate(d *schema.ResourceData, m interface{}, createStage func() stage) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	s := createStage()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &s); err != nil {
		return err
	}
	stage := s.(stage)
	stage.SetRefID(d.Id())

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	cs, err := stage.toClientStage(&m.(*Services).Config)
	if err != nil {
		return err
	}
	err = pipeline.UpdateStage(cs)
	if err != nil {
		return err
	}

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated pipeline stages:", d.Id())
	return resourcePipelineStageRead(d, m, createStage)
}

func resourcePipelineStageDelete(d *schema.ResourceData, m interface{}, createStage func() stage) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	// TODO duplicated code from resourcePipelineStageUpdate
	s := createStage()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &s); err != nil {
		return err
	}
	stage := s.(stage)
	stage.SetRefID(d.Id())

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	err = pipeline.DeleteStage(stage.GetRefID())
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
