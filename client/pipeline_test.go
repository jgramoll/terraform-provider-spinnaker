package client

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestGetApplicationPipelines(t *testing.T) {
	_, err := client.GetApplicationPipelines("career")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPipeline(t *testing.T) {
	pipeline, err := client.GetPipeline("career", "Bridge Nav Edge")
	if err != nil {
		t.Fatal(err)
	}

	if pipeline.Name != "Bridge Nav Edge" {
		t.Fatal("should be pipeline Bridge Nav Edge")
	}
}

func TestGetPipelineByID(t *testing.T) {
	pipeline, err := client.GetPipelineByID("13caa723-114a-4d05-94f0-7f786f981c10")
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
	err := client.CreatePipeline(&CreatePipelineRequest{
		Name:        name,
		Application: app,
	})
	if err != nil {
		t.Fatal(err)
	}

	var pipeline *Pipeline
	pipeline, err = client.GetPipeline(app, name)
	if err != nil {
		t.Fatal(err)
	}

	newApp := "career"
	newName := fmt.Sprintf("My New Name Pipe %d", rand.Int())
	pipeline.Name = newName
	pipeline.Application = newApp
	err = client.UpdatePipeline(pipeline)
	if err != nil {
		t.Fatal(err)
	}

	pipeline, err = client.GetPipeline(newApp, newName)
	if err != nil {
		t.Fatal(err)
	}
	if pipeline.Application != newApp {
		t.Fatalf("app should now be %s, not %s", newApp, pipeline.Application)
	}
	if pipeline.Name != newName {
		t.Fatalf("name should now be %s, not %s", newName, pipeline.Name)
	}

	err = client.DeletePipeline(pipeline)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCleanup(t *testing.T) {
	pipelines, err := client.GetApplicationPipelines("app")
	if err != nil {
		t.Fatal(err)
	}

	for _, pipe := range *pipelines {
		if strings.Contains(pipe.Name, "tf-acc") {
			client.DeletePipeline(pipe)
		}
		if strings.Contains(pipe.Name, "My Test") {
			client.DeletePipeline(pipe)
		}
	}
}
