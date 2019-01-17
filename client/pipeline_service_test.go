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
	applicationName = "career"
}

func TestGetApplicationPipelines(t *testing.T) {
	pipelines, err := pipelineService.GetApplicationPipelines(applicationName)
	if err != nil {
		t.Fatal(err)
	}
	if len(*pipelines) == 0 {
		t.Fatal("no pipelines")
	}
}

func TestGetPipeline(t *testing.T) {
	pipelineName := "Bridge Nav Edge"
	pipeline, err := pipelineService.GetPipeline(applicationName, pipelineName)
	if err != nil {
		t.Fatal(err)
	}

	if pipeline.Name != pipelineName {
		t.Fatalf("should be pipeline %s, not %s", pipelineName, pipeline.Name)
	}
}

func TestGetPipelineByID(t *testing.T) {
	pipelineID := "13caa723-114a-4d05-94f0-7f786f981c10"
	pipeline, err := pipelineService.GetPipelineByID(pipelineID)
	if err != nil {
		t.Fatal(err)
	}

	if pipeline.ID != pipelineID {
		t.Fatalf("should be pipeline id %s, not %s", pipelineID, pipeline.ID)
	}
	if pipeline.Name != "test" {
		t.Fatalf("should be pipeline test, not %s", pipeline.Name)
	}
}

func TestCreateUpdateDeletePipeline(t *testing.T) {
	name := fmt.Sprintf("My Test Pipe %d", rand.Int())
	err := pipelineService.CreatePipeline(&CreatePipelineRequest{
		Name:        name,
		Application: applicationName,
	})
	if err != nil {
		t.Fatal(err)
	}

	var pipeline *Pipeline
	pipeline, err = pipelineService.GetPipeline(applicationName, name)
	if err != nil {
		t.Fatal(err)
	}

	newApp := "app"
	newName := fmt.Sprintf("My New Name Pipe %d", rand.Int())
	pipeline.Name = newName
	pipeline.Application = newApp
	pipeline.Stages = &[]Stage{
		NewBakeStage(),
	}
	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		t.Fatal(err)
	}

	var updatedPipeline *Pipeline
	updatedPipeline, err = pipelineService.GetPipeline(newApp, newName)
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
