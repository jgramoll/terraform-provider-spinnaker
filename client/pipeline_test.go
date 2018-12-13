package client

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
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

func TestCreatePipeline(t *testing.T) {
	pipeline := Pipeline{
		Name:        fmt.Sprintf("My Test Pipe %d", rand.Int()),
		Application: "app",
	}
	err := client.CreatePipeline(&pipeline)
	if err != nil {
		t.Fatal(err)
	}

	err = client.DeletePipeline(&pipeline)
	if err != nil {
		t.Fatal(err)
	}
}
