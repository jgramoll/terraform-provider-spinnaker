package client

import (
	"fmt"
)

type Pipeline struct {
	Application          string
	Disabled             bool
	Id                   string
	Index                int
	KeepWaitingPipelines bool
	LastModifiedBy       string
	LimitConcurrent      bool
	Name                 string
	// notifications    []Notification
	// parameterConfig  []
	// Stages   []Stage
	// Triggers []Trigger
	UpdateTs string
}

func (client *Client) GetPipeline(applicationName string, pipelineName string) (*Pipeline, error) {
	// TODO this is return list of pipeline
	path := fmt.Sprintf("/applications/%s/pipelineConfigs", applicationName)
	req, err := client.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var pipelines []Pipeline
	_, respErr := client.Do(req, &pipelines)
	if respErr != nil {
		return nil, respErr
	}

	for _, pipeline := range pipelines {
		if pipeline.Name == pipelineName {
			return &pipeline, nil
		}
	}

	return nil, fmt.Errorf("Pipeline %s not found", pipelineName)
}
