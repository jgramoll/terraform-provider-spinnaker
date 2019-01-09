package client

import (
	"testing"
)

var pipelineStage PipelineStage

func init() {
	pipelineStage = *NewPipelineStage()
}

func TestPipelineStageGetName(t *testing.T) {
	name := "New Pipeline Stage"
	pipelineStage.Name = name
	if pipelineStage.GetName() != name {
		t.Fatalf("Pipeline stage GetName() should be %s, not \"%s\"", name, pipelineStage.GetName())
	}
}

func TestPipelineStageGetType(t *testing.T) {
	if pipelineStage.GetType() != PipelineType {
		t.Fatalf("Pipeline stage GetType() should be %s, not \"%s\"", PipelineType, pipelineStage.GetType())
	}
	if pipelineStage.Type != PipelineType {
		t.Fatalf("Pipeline stage Type should be %s, not \"%s\"", PipelineType, pipelineStage.Type)
	}
}
