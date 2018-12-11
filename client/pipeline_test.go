package client

import (
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
