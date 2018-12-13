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
	path := fmt.Sprintf("/applications/%s/pipelineConfigs/%s", applicationName, pipelineName)
	req, err := client.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var pipeline Pipeline
	_, respErr := client.DoWithResponse(req, &pipeline)
	if respErr != nil {
		return nil, respErr
	}

	return &pipeline, nil
}

// CreatePipeline in application
// Example: {"name":"test","stages":[],"triggers":[],"application":"career","limitConcurrent":true,"keepWaitingPipelines":false,"index":5}
func (client *Client) CreatePipeline(pipeline *Pipeline) error {
	// TODO is there a way to create map from object
	data := map[string]interface{}{
		"application": pipeline.Application,
		"name":        pipeline.Name,
	}

	path := "/pipelines"
	req, err := client.NewRequestWithBody("POST", path, data)
	if err != nil {
		return err
	}

	_, respErr := client.Do(req)
	return respErr
}

// DeletePipeline in application
func (client *Client) DeletePipeline(pipeline *Pipeline) error {
	path := fmt.Sprintf("/pipelines/%s/%s", pipeline.Application, pipeline.Name)
	req, err := client.NewRequest("DELETE", path)
	if err != nil {
		return err
	}

	_, respErr := client.Do(req)
	return respErr
}
