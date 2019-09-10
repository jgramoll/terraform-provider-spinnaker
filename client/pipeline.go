package client

import (
	"github.com/mitchellh/mapstructure"
)

// SerializablePipeline deploy pipeline in application
type SerializablePipeline struct {
	Application          string                 `json:"application"`
	AppConfig            map[string]interface{} `json:"appConfig"`
	Disabled             bool                   `json:"disabled"`
	ID                   string                 `json:"id"`
	Index                int                    `json:"index"`
	KeepWaitingPipelines bool                   `json:"keepWaitingPipelines"`
	LimitConcurrent      bool                   `json:"limitConcurrent"`
	Name                 string                 `json:"name"`
	Roles                *[]string              `json:"roles"`
	ServiceAccount       string                 `json:"serviceAccount,omitempty"`

	Locked          *Locked               `json:"locked,omitempty"`
	ParameterConfig *[]*PipelineParameter `json:"parameterConfig"`
	Triggers        []*Trigger            `json:"triggers"`
}

// Pipeline deploy pipeline in application
type Pipeline struct {
	SerializablePipeline

	Notifications *[]*Notification `json:"notifications"`
	Stages        *[]Stage         `json:"stages"`
}

// Locked enable/disable to edit Pipeline over the UI
type Locked struct {
	UI            bool `json:"ui"`
	AllowUnlockUI bool `json:"allowUnlockUi"`
}

// NewSerializablePipeline Pipeline with default values
func NewSerializablePipeline() SerializablePipeline {
	return SerializablePipeline{
		Disabled:             false,
		KeepWaitingPipelines: false,
		LimitConcurrent:      true,
		Triggers:             []*Trigger{},
		AppConfig:            map[string]interface{}{},
		ParameterConfig:      &[]*PipelineParameter{},
	}
}

// NewPipeline Pipeline with default values
func NewPipeline() *Pipeline {
	return &Pipeline{
		SerializablePipeline: NewSerializablePipeline(),
	}
}

func parsePipeline(pipelineHash map[string]interface{}) (*Pipeline, error) {
	serializablePipeline := NewSerializablePipeline()
	if err := mapstructure.Decode(pipelineHash, &serializablePipeline); err != nil {
		return nil, err
	}

	stages, err := parseStages(pipelineHash["stages"])
	if err != nil {
		return nil, err
	}

	notifications, err := parseNotifications(pipelineHash["notifications"])
	if err != nil {
		return nil, err
	}

	locked, err := parseLocked(pipelineHash["locked"])
	if err != nil {
		return nil, err
	}

	if locked != nil {
		serializablePipeline.Locked = locked
	}

	return &Pipeline{
		SerializablePipeline: serializablePipeline,
		Notifications:        notifications,
		Stages:               stages,
	}, nil
}

func parseLocked(lockedHash interface{}) (*Locked, error) {
	if lockedHash == nil {
		return nil, nil
	}

	lockedToParse := lockedHash.(map[string]interface{})
	return &Locked{
		UI:            lockedToParse["ui"].(bool),
		AllowUnlockUI: lockedToParse["allowUnlockUi"].(bool),
	}, nil
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
