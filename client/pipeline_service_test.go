// +build integration

package client

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var pipelineService *PipelineService
var applicationName string

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	pipelineService = &PipelineService{newTestClient()}
	applicationName = "app"
}

func createPipeline(t *testing.T) *CreatePipelineRequest {
	name := fmt.Sprintf("My Test Pipe %d", rand.Int())
	pipeline := CreatePipelineRequest{
		Name:        name,
		Application: applicationName,
	}
	err := pipelineService.CreatePipeline(&pipeline)
	if err != nil {
		t.Fatal(err)
	}
	return &pipeline
}

func TestGetApplicationPipelines(t *testing.T) {
	createPipeline(t)
	createPipeline(t)

	pipelines, err := pipelineService.GetApplicationPipelines(applicationName)
	if err != nil {
		t.Fatal(err)
	}
	if len(*pipelines) < 2 {
		t.Fatalf("should be at least 2 pipelines not %v", len(*pipelines))
	}
}

func TestGetPipeline(t *testing.T) {
	pipelineReq := createPipeline(t)
	pipeline, err := pipelineService.GetPipeline(pipelineReq.Application, pipelineReq.Name)
	if err != nil {
		t.Fatal(err)
	}

	if pipeline.Name != pipelineReq.Name {
		t.Fatalf("should be pipeline %s, not %s", pipelineReq.Name, pipeline.Name)
	}

	pipelineID := pipeline.ID
	pipeline, err = pipelineService.GetPipelineByID(pipelineID)
	if err != nil {
		t.Fatal(err)
	}

	if pipeline.ID != pipelineID {
		t.Fatalf("should be pipeline id %s, not %s", pipelineID, pipeline.ID)
	}
	if pipeline.Name != pipelineReq.Name {
		t.Fatalf("should be pipeline %s, not %s", pipelineReq.Name, pipeline.Name)
	}
}

func TestCreateUpdateDeletePipeline(t *testing.T) {
	pipelineReq := createPipeline(t)

	var pipeline *Pipeline
	pipeline, err := pipelineService.GetPipeline(pipelineReq.Application, pipelineReq.Name)
	if err != nil {
		t.Fatal(err)
	}

	newName := fmt.Sprintf("My New Name Pipe %d", rand.Int())
	pipeline.Name = newName
	pipeline.Stages = &[]Stage{
		NewBakeStage(),
	}
	pipeline.ParameterConfig = &[]*PipelineParameter{
		{Name: "new parameter"},
		{Name: "descriptive parameter", Description: "This is a very descriptive parameter."},
		{
			Name:        "Options parameter",
			Default:     "mosdef",
			Description: "Setting parameter via options.",
			HasOptions:  true,
			Label:       "mosdef label",
			Options: &[]*PipelineParameterOption{
				{Value: "something"},
			},
			Required: true,
		},
	}
	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		t.Fatal(err)
	}

	var updatedPipeline *Pipeline
	updatedPipeline, err = pipelineService.GetPipeline(pipeline.Application, newName)
	if err != nil {
		t.Fatal(err)
	}

	// Not sure why spinnaker is playing with indexes...
	pipeline.Index = updatedPipeline.Index
	err = pipeline.equals(updatedPipeline)
	if err != nil {
		t.Fatal(err)
	}

	err = pipelineService.DeletePipeline(pipeline)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCleanup(t *testing.T) {
	pipelines, err := pipelineService.GetApplicationPipelines(applicationName)
	if err != nil {
		t.Fatal(err)
	}

	for _, pipe := range *pipelines {
		if strings.Contains(pipe.Name, "tf-acc") {
			pipelineService.DeletePipeline(pipe)
		}
		if strings.Contains(pipe.Name, "My Test") {
			pipelineService.DeletePipeline(pipe)
		}
		if strings.Contains(pipe.Name, "My New Name") {
			pipelineService.DeletePipeline(pipe)
		}
	}
}
