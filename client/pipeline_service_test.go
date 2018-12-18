package client

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var pipelineService *PipelineService

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	pipelineService = &PipelineService{newTestClient()}
}

func TestGetApplicationPipelines(t *testing.T) {
	pipelines, err := pipelineService.GetApplicationPipelines("career")
	if err != nil {
		t.Fatal(err)
	}
	if len(*pipelines) == 0 {
		t.Fatal("no pipelines")
	}
}

func TestGetPipeline(t *testing.T) {
	pipeline, err := pipelineService.GetPipeline("career", "Bridge Nav Edge")
	if err != nil {
		t.Fatal(err)
	}

	if pipeline.Name != "Bridge Nav Edge" {
		t.Fatal("should be pipeline Bridge Nav Edge")
	}
}

func TestGetPipelineByID(t *testing.T) {
	pipeline, err := pipelineService.GetPipelineByID("13caa723-114a-4d05-94f0-7f786f981c10")
	if err != nil {
		t.Fatal(err)
	}

	if pipeline.Name != "test" {
		t.Fatal("should be pipeline test")
	}
}

func TestCreateUpdateDeletePipeline(t *testing.T) {
	name := fmt.Sprintf("My Test Pipe %d", rand.Int())
	app := "app"
	err := pipelineService.CreatePipeline(&CreatePipelineRequest{
		Name:        name,
		Application: app,
	})
	if err != nil {
		t.Fatal(err)
	}

	var pipeline *Pipeline
	pipeline, err = pipelineService.GetPipeline(app, name)
	if err != nil {
		t.Fatal(err)
	}

	newApp := "career"
	newName := fmt.Sprintf("My New Name Pipe %d", rand.Int())
	pipeline.Name = newName
	pipeline.Application = newApp
	pipeline.Stages = []Stage{
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
	pipelines, err := pipelineService.GetApplicationPipelines("career")
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
