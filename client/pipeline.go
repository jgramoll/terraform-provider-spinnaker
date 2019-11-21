package client

import (
	"github.com/mitchellh/mapstructure"
)

// Pipeline deploy pipeline in application
type Pipeline struct {
	Application          string                 `json:"application"`
	AppConfig            map[string]interface{} `json:"appConfig"`
	Disabled             bool                   `json:"disabled"`
	ID                   string                 `json:"id"`
	Index                int                    `json:"index"`
	KeepWaitingPipelines bool                   `json:"keepWaitingPipelines"`
	LimitConcurrent      bool                   `json:"limitConcurrent"`
	Name                 string                 `json:"name"`
	ParameterConfig      *[]*PipelineParameter  `json:"parameterConfig"`
	Roles                *[]string              `json:"roles"`
	ServiceAccount       string                 `json:"serviceAccount,omitempty"`
	Triggers             []Trigger              `json:"triggers"`

	Notifications *[]*Notification `json:"notifications"`
	Stages        *[]Stage         `json:"stages"`
}

// NewPipeline Pipeline with default values
func NewPipeline() *Pipeline {
	return &Pipeline{
		Disabled:             false,
		KeepWaitingPipelines: false,
		LimitConcurrent:      true,
		Triggers:             []Trigger{},
		AppConfig:            map[string]interface{}{},
		ParameterConfig:      &[]*PipelineParameter{},
	}
}

func parsePipeline(pipelineHash map[string]interface{}) (*Pipeline, error) {
	pipeline := NewPipeline()

	stages, err := parseStages(pipelineHash["stages"])
	if err != nil {
		return nil, err
	}
	pipeline.Stages = stages
	delete(pipelineHash, "stages")

	notifications, err := parseNotifications(pipelineHash["notifications"])
	if err != nil {
		return nil, err
	}
	pipeline.Notifications = notifications
	delete(pipelineHash, "notifications")

	triggers, err := parseTriggers(pipelineHash["triggers"])
	if err != nil {
		return nil, err
	}
	pipeline.Triggers = *triggers
	delete(pipelineHash, "triggers")

	if err := mapstructure.Decode(pipelineHash, pipeline); err != nil {
		return nil, err
	}
	return pipeline, nil
}

// AppendStage append stage
func (pipeline *Pipeline) AppendStage(stage Stage) {
	if pipeline.Stages == nil {
		pipeline.Stages = &[]Stage{}
	}
	stages := append(*pipeline.Stages, stage)
	pipeline.Stages = &stages
}

// AppendNotification append notification
func (pipeline *Pipeline) AppendNotification(notification *Notification) {
	if pipeline.Notifications == nil {
		pipeline.Notifications = &[]*Notification{}
	}
	newNotifications := append(*pipeline.Notifications, notification)
	pipeline.Notifications = &newNotifications
}
