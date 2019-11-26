package client

import (
	"testing"
)

var scaleManifestStage ScaleManifestStage

func init() {
	scaleManifestStage = *NewScaleManifestStage()
}

func TestScaleManifestStageGetName(t *testing.T) {
	name := "New Scale Manifest"
	scaleManifestStage.Name = name
	if scaleManifestStage.GetName() != name {
		t.Fatalf("Scale Manifest stage GetName() should be %s, not \"%s\"", name, scaleManifestStage.GetName())
	}
}

func TestScaleManifestStageGetType(t *testing.T) {
	if scaleManifestStage.GetType() != ScaleManifestStageType {
		t.Fatalf("Scale Manifest stage GetType() should be %s, not \"%s\"", ScaleManifestStageType, scaleManifestStage.GetType())
	}
	if scaleManifestStage.Type != ScaleManifestStageType {
		t.Fatalf("Scale Manifest stage Type should be %s, not \"%s\"", ScaleManifestStageType, scaleManifestStage.Type)
	}
}
