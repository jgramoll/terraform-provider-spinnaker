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

type stage interface {
	toClientStage() client.Stage
	SetResourceData()
}

type baseStage struct {
	Name  string           `mapstructure:"name"`
	RefID string           `mapstructure:"ref_id"`
	Type  client.StageType `mapstructure:"type"`
}

func resourcePipelineStageCreate(d *schema.ResourceData, m interface{}, createStage func() stage) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	stage := createStage()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &stage); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	// stage.RefID = id.String()

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
	// return resourcePipelineStageRead(d, m, stageType)
	return nil
}

func resourcePipelineStageRead(d *schema.ResourceData, m interface{}, stageType string) error {
	pipelineID := d.Get("pipeline").(string)
	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(pipelineID)
	if err != nil {
		log.Println("[WARN] No Pipeline found:", err)
		d.SetId("")
		return nil
	}

	var stage client.Stage
	stage, err = getStage(pipeline.Stages, d.Id())
	if err != nil {
		log.Println("[WARN] No Pipeline Stage found:", err)
		d.SetId("")
	} else {
		// TODO
		d.Set("name", stage.GetName())
		// d.Set("vm_type", stage.VMType)
	}

	return nil
}

func getStage(stages []client.Stage, stageID string) (client.Stage, error) {
	// for _, stage := range stages {
	// if stage.RefID == stageID {
	// 	return &stage, nil
	// }
	// }
	return nil, ErrStageNotFound
}

func resourcePipelineStageUpdate(d *schema.ResourceData, m interface{}, stageType string) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	var stage stage
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &stage); err != nil {
		return err
	}
	// stage.RefID = d.Id()
	// stage.Type = stageType

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get("pipeline").(string))
	if err != nil {
		return err
	}

	// err = updateStages(pipeline, client.Stage(stage))
	if err != nil {
		return err
	}

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated pipeline stages:", d.Id())
	return resourcePipelineStageRead(d, m, stageType)
}

func updateStages(pipeline *client.Pipeline, stage client.Stage) error {
	// for i, t := range pipeline.Stages {
	// if t.RefID == stage.RefID {
	// 	pipeline.Stages[i] = stage
	// 	return nil
	// }
	// }
	// return ErrStageNotFound
	return nil
}

func resourcePipelineStageDelete(d *schema.ResourceData, m interface{}, stageType string) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	var stage stage
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &stage); err != nil {
		return err
	}
	// stage.RefID = d.Id()
	// stage.Type = stageType

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get("pipeline").(string))
	if err != nil {
		return err
	}

	// err = deleteStage(pipeline, client.Stage(stage))
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

func deleteStage(pipeline *client.Pipeline, stage client.Stage) error {
	// for i, t := range pipeline.Stages {
	// 	if t.RefID == stage.RefID {
	// 		pipeline.Stages = append(pipeline.Stages[:i], pipeline.Stages[i+1:]...)
	// 		return nil
	// 	}
	// }
	// return ErrStageNotFound
	return nil
}
