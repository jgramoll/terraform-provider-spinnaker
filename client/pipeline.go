package client

import (
	"fmt"
)

// UpdatePipelineRequest used to create pipeline
type CreatePipelineRequest struct {
	Application string `json:"application"`
	Name        string `json:"name"`
}

// Pipeline deploy pipeline in application
type Pipeline struct {
	Application          string `json:"application"`
	Disabled             bool   `json:"disabled"`
	ID                   string `json:"id"`
	Index                int    `json:"index"`
	KeepWaitingPipelines bool   `json:"keepWaitingPipelines"`
	// LastModifiedBy       string `json:"lastModifiedBy"`
	LimitConcurrent bool   `json:"limitConcurrent"`
	Name            string `json:"name"`
	// notifications    []Notification
	// parameterConfig  []
	// Stages   []Stage
	// Triggers []Trigger
	// UpdateTs string `json:"updateTs"`
}

// NewPipeline Pipeline with default values
func NewPipeline() *Pipeline {
	return &Pipeline{
		Disabled:             false,
		KeepWaitingPipelines: false,
		LimitConcurrent:      true,
	}
}

// GetApplicationPipelines get all pipelines for an application
func (client *Client) GetApplicationPipelines(applicationName string) (*[]*Pipeline, error) {
	path := fmt.Sprintf("/applications/%s/pipelineConfigs", applicationName)
	req, err := client.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var pipelines []*Pipeline
	_, respErr := client.DoWithResponse(req, &pipelines)
	if respErr != nil {
		return nil, respErr
	}

	return &pipelines, nil
}

// GetPipeline get pipeline by name and application
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
func (client *Client) CreatePipeline(pipeline *CreatePipelineRequest) error {
	path := "/pipelines"
	req, err := client.NewRequestWithBody("POST", path, pipeline)
	if err != nil {
		return err
	}

	_, respErr := client.Do(req)
	return respErr
}

// UpdatePipeline in application
func (client *Client) UpdatePipeline(pipeline *Pipeline) error {
	path := "/pipelines"
	req, err := client.NewRequestWithBody("POST", path, pipeline)
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
