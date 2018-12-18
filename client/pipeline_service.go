package client

import (
	"errors"
	"fmt"
	"net/http"
)

// ErrPipelineNotFound pipeline not found
var ErrPipelineNotFound = errors.New("Could not find pipeline")

// PipelineService used to manage pipelines
type PipelineService struct {
	*Client
}

// CreatePipelineRequest used to create pipeline
type CreatePipelineRequest struct {
	Application string `json:"application"`
	Name        string `json:"name"`
}

// GetApplicationPipelines get all pipelines for an application
func (service *PipelineService) GetApplicationPipelines(applicationName string) (*[]*Pipeline, error) {
	path := fmt.Sprintf("/applications/%s/pipelineConfigs", applicationName)
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	return service.parsePipelinesRequest(req)
}

// GetPipelineByID get pipeline by id
func (service *PipelineService) GetPipelineByID(id string) (*Pipeline, error) {
	path := fmt.Sprintf("/pipelineConfigs/%s/history?limit=1", id)
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var pipelines *[]*Pipeline
	pipelines, err = service.parsePipelinesRequest(req)

	if len(*pipelines) == 0 {
		return nil, ErrPipelineNotFound
	}

	return (*pipelines)[0], nil
}

// GetPipeline get pipeline by name and application
func (service *PipelineService) GetPipeline(applicationName string, pipelineName string) (*Pipeline, error) {
	path := fmt.Sprintf("/applications/%s/pipelineConfigs/%s", applicationName, pipelineName)
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var pipelineHash map[string]interface{}
	_, respErr := service.DoWithResponse(req, &pipelineHash)
	if respErr != nil {
		return nil, respErr
	}

	return parsePipeline(pipelineHash)
}

// CreatePipeline in application
func (service *PipelineService) CreatePipeline(pipeline *CreatePipelineRequest) error {
	path := "/pipelines"
	req, err := service.NewRequestWithBody("POST", path, pipeline)
	if err != nil {
		return err
	}

	_, respErr := service.Do(req)
	return respErr
}

// UpdatePipeline in application
func (service *PipelineService) UpdatePipeline(pipeline *Pipeline) error {
	path := "/pipelines"
	req, err := service.NewRequestWithBody("POST", path, pipeline)
	if err != nil {
		return err
	}

	_, respErr := service.Do(req)
	return respErr
}

// DeletePipeline in application
func (service *PipelineService) DeletePipeline(pipeline *Pipeline) error {
	path := fmt.Sprintf("/pipelines/%s/%s", pipeline.Application, pipeline.Name)
	req, err := service.NewRequest("DELETE", path)
	if err != nil {
		return err
	}

	_, respErr := service.Do(req)
	return respErr
}

func (service *PipelineService) parsePipelinesRequest(req *http.Request) (*[]*Pipeline, error) {
	var pipelinesHash []map[string]interface{}
	_, respErr := service.DoWithResponse(req, &pipelinesHash)
	if respErr != nil {
		return nil, respErr
	}

	var pipelines []*Pipeline
	for _, pipelineHash := range pipelinesHash {
		pipeline, err := parsePipeline(pipelineHash)
		if err != nil {
			return nil, err
		}
		pipelines = append(pipelines, pipeline)
	}
	return &pipelines, nil
}
