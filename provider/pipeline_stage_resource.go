package provider

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

// ErrStageNotFound stage not found
var ErrStageNotFound = errors.New("Could not find stage")

func resourcePipelineStageCreate(d *schema.ResourceData, m interface{}, createStage func() interface{}) error {
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
	pipeline, err := pipelineService.GetPipelineByID(d.Get("pipeline").(string))
	if err != nil {
		return err
	}

	pipeline.Stages = append(pipeline.Stages, stage.toClientStage())
	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Creating pipeline stage:", id)
	d.SetId(id.String())
	return resourcePipelineStageRead(d, m, createStage)
}

func resourcePipelineStageRead(d *schema.ResourceData, m interface{}, createStage func() interface{}) error {
	pipelineID := d.Get("pipeline").(string)
	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(pipelineID)
	if err != nil {
		log.Println("[WARN] No Pipeline found:", err)
		d.SetId("")
		return nil
	}

	var cStage client.Stage
	cStage, err = getStage(pipeline.Stages, d.Id())
	if err != nil {
		log.Println("[WARN] No Pipeline Stage found:", err)
		d.SetId("")
	} else {
		s := createStage().(stage)
		s = s.fromClientStage(cStage)
		log.Println("[INFO] Updating from stage", cStage)
		s.SetResourceData(d)
	}

	return nil
}

func resourcePipelineStageUpdate(d *schema.ResourceData, m interface{}, createStage func() interface{}) error {
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
	pipeline, err := pipelineService.GetPipelineByID(d.Get("pipeline").(string))
	if err != nil {
		return err
	}

	err = updateStages(pipeline, stage.toClientStage())
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

func resourcePipelineStageDelete(d *schema.ResourceData, m interface{}, createStage func() interface{}) error {
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
	pipeline, err := pipelineService.GetPipelineByID(d.Get("pipeline").(string))
	if err != nil {
		return err
	}

	err = deleteStage(pipeline, stage.toClientStage())
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
