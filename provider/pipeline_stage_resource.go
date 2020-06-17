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

func stageResource(in map[string]*schema.Schema) map[string]*schema.Schema {
	out := map[string]*schema.Schema{
		PipelineKey: {
			Type:        schema.TypeString,
			Description: "Id of the pipeline",
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "Name of the stage",
			Required:    true,
		},
		"requisite_stage_ref_ids": {
			Type:        schema.TypeList,
			Description: "Stage(s) that must be complete before this one",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"expected_artifact": {
			Type:        schema.TypeList,
			Description: "Expected artifacts for stage",
			Optional:    true,
			Elem:        expectedArtifactResource(),
		},
		"notification": {
			Type:        schema.TypeList,
			Description: "Notifications to send for stage results",
			Optional:    true,
			Elem:        notificationResource(),
		},
		"complete_other_branches_then_fail": {
			Type:        schema.TypeBool,
			Description: "halt this branch and fail the pipeline once other branches complete. Prevents any stages that depend on this stage from running, but allows other branches of the pipeline to run. The pipeline will be marked as failed once complete.",
			Optional:    true,
			Default:     false,
		},
		"continue_pipeline": {
			Type:        schema.TypeBool,
			Description: "If false, marks the stage as successful right away without waiting for the jenkins job to complete",
			Optional:    true,
			Default:     false,
		},
		"fail_pipeline": {
			Type:        schema.TypeBool,
			Description: "If the stage fails, immediately halt execution of all running stages and fails the entire execution",
			Optional:    true,
			Default:     true,
		},
		"fail_on_failed_expressions": {
			Type:        schema.TypeBool,
			Description: "The stage will be marked as failed if it contains any failed expressions",
			Optional:    true,
			Default:     false,
		},
		"override_timeout": {
			Type:        schema.TypeBool,
			Description: "[Deprecated, use stage_timeout_ms] Allows you to override the amount of time the stage can run before failing.\nNote: this represents the overall time the stage has to complete (the sum of all the task times).",
			Optional:    true,
			Default:     false,
		},
		"stage_timeout_ms": {
			Type:        schema.TypeInt,
			Description: "Allows you to declare the amount of time the stage can run before failing, if override timeout is enabled.\nNote: this represents the overall time the stage has to complete (the sum of all the task times).",
			Optional:    true,
		},
		"restrict_execution_during_time_window": {
			Type:        schema.TypeBool,
			Description: "Restrict execution to specific time windows",
			Optional:    true,
			Default:     false,
		},
		"restricted_execution_window": {
			Type:        schema.TypeList,
			Description: "Time windows to restrict execution",
			Optional:    true,
			MaxItems:    1,
			Elem:        restrictedExecutionWindowResource(),
		},
		"stage_enabled": {
			Type:        schema.TypeList,
			Description: "Stage will only execute when the supplied expression evaluates true.\nThe expression does not need to be wrapped in ${ and }.\nIf this expression evaluates to false, the stages following this stage will still execute.",
			Optional:    true,
			MaxItems:    1,
			Elem:        stageEnabledResource(),
		},
	}

	// merge input
	for k, v := range in {
		out[k] = v
	}

	return out
}

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

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	cs, err := s.toClientStage(m.(*Services).Config, id.String())
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
	if err == client.ErrStageNotFound {
		log.Println("[WARN] No Pipeline Stage found")
		d.SetId("")
		return nil
	} else if err != nil {
		log.Println("[ERROR] Error on get Pipeline stage:", err)
		d.SetId("")
		return err
	}

	s := createStage().(stage)
	s, err = s.fromClientStage(cStage)
	if err != nil {
		log.Println("[ERROR] Error on reading pipeline stage:", err)
		d.SetId("")
		return err
	}
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

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	cs, err := s.toClientStage(m.(*Services).Config, d.Id())
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

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	err = pipeline.DeleteStage(d.Id())
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
