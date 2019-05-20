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
	if err != nil {
		return nil, err
	}

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
	_, err = service.DoWithResponse(req, &pipelineHash)
	if err != nil {
		return nil, err
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

	_, err = service.Do(req)
	return err
}

// UpdatePipeline in application
func (service *PipelineService) UpdatePipeline(pipeline *Pipeline) error {
	path := "/pipelines"
	// Hack around async updates to the pipeline
	// If we don't do this we get periodic 400s
	_, err := service.DoWithRetry(func() (*http.Request, error) {
		return service.NewRequestWithBody("POST", path, pipeline)
	})
	return err
}

// DeletePipeline in application
func (service *PipelineService) DeletePipeline(pipeline *Pipeline) error {
	path := fmt.Sprintf("/pipelines/%s/%s", pipeline.Application, pipeline.Name)
	req, err := service.NewRequest("DELETE", path)
	if err != nil {
		return err
	}

	_, err = service.Do(req)
	return err
}

func (service *PipelineService) parsePipelinesRequest(req *http.Request) (*[]*Pipeline, error) {
	var pipelinesHash []map[string]interface{}
	_, err := service.DoWithResponse(req, &pipelinesHash)
	if err != nil {
		return nil, err
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
