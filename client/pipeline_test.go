package client

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGetPipeline(t *testing.T) {
	pipeline, err := client.GetPipeline("career", "Bridge Nav Edge")
	if err != nil {
		t.Fatal(err)
	}

	if pipeline.Name != "Bridge Nav Edge" {
		t.Fatal("should be pipeline Bridge Nav Edge")
	}
}

func TestPostPipeline(t *testing.T) {
	name := fmt.Sprintf("My Test Pipe %d", rand.Int())
	err := client.PostPipeline("career", name)
	if err != nil {
		t.Fatal(err)
	}
}
