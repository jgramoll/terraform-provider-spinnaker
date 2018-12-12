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
	_, respErr := client.Do(req, &pipeline)
	if respErr != nil {
		return nil, respErr
	}

	return &pipeline, nil
}

// {"name":"test","stages":[],"triggers":[],"application":"career","limitConcurrent":true,"keepWaitingPipelines":false,"index":5}
func (client *Client) PostPipeline(applicationName string, pipelineName string) error {
	data := map[string]interface{}{
		"application": applicationName,
		"name":        pipelineName,
	}

	path := "/pipelines"
	_, err := client.NewRequestWithBody("POST", path, data)
	if err != nil {
		return err
	}
	return nil
}
