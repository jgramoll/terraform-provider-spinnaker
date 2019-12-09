package client

import (
	"testing"
)

var findArtifactsFromResourceStage FindArtifactsFromResourceStage

func init() {
	findArtifactsFromResourceStage = *NewFindArtifactsFromResourceStage()
}

func TestNewFindArtifactsFromResourceStage(t *testing.T) {
	if findArtifactsFromResourceStage.Type != FindArtifactsFromResourceStageType {
		t.Fatalf("Deploy stage type should be %s, not \"%s\"", FindArtifactsFromResourceStageType, findArtifactsFromResourceStage.Type)
	}
}

func TestFindArtifactsFromResourceStageGetName(t *testing.T) {
	name := "New Deploy"
	findArtifactsFromResourceStage.Name = name
	if findArtifactsFromResourceStage.GetName() != name {
		t.Fatalf("Deploy stage GetName() should be %s, not \"%s\"", name, findArtifactsFromResourceStage.GetName())
	}
}

func TestFindArtifactsFromResourceStageGetType(t *testing.T) {
	if findArtifactsFromResourceStage.GetType() != FindArtifactsFromResourceStageType {
		t.Fatalf("Deploy stage GetType() should be %s, not \"%s\"", FindArtifactsFromResourceStageType, findArtifactsFromResourceStage.GetType())
	}
	if findArtifactsFromResourceStage.Type != FindArtifactsFromResourceStageType {
		t.Fatalf("Deploy stage Type should be %s, not \"%s\"", FindArtifactsFromResourceStageType, findArtifactsFromResourceStage.Type)
	}
}
