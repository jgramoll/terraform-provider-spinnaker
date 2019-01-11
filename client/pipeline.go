package client

import (
	"github.com/mitchellh/mapstructure"
)

// ParameterConfig config for pipeline parameters
type ParameterConfig struct{}

// SerializablePipeline deploy pipeline in application
type SerializablePipeline struct {
	Application          string            `json:"application"`
	Disabled             bool              `json:"disabled"`
	ID                   string            `json:"id"`
	Index                int               `json:"index"`
	KeepWaitingPipelines bool              `json:"keepWaitingPipelines"`
	LimitConcurrent      bool              `json:"limitConcurrent"`
	Name                 string            `json:"name"`
	ParameterConfig      []ParameterConfig `json:"parameterConfig"`
	Triggers             []Trigger         `json:"triggers"`
	// TODO pointers?
}

// Pipeline deploy pipeline in application
type Pipeline struct {
	SerializablePipeline

	Notifications *[]Notification `json:"notifications"`
	Stages        *[]Stage        `json:"stages"`
}

// NewSerializablePipeline Pipeline with default values
func NewSerializablePipeline() SerializablePipeline {
	return SerializablePipeline{
		Disabled:             false,
		KeepWaitingPipelines: false,
		LimitConcurrent:      true,
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

	return &Pipeline{
		SerializablePipeline: serializablePipeline,
		Notifications:        notifications,
		Stages:               stages,
	}, nil
}
