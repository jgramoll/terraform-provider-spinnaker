package client

import (
	"testing"
)

var enableManifestStage EnableManifestStage

func init() {
	enableManifestStage = *NewEnableManifestStage()
}

func TestEnableManifestStageGetName(t *testing.T) {
	name := "New Enable Manifest"
	enableManifestStage.Name = name
	if enableManifestStage.GetName() != name {
		t.Fatalf("Delete Manifest stage GetName() should be %s, not \"%s\"", name, enableManifestStage.GetName())
	}
}

func TestEnableManifestStageGetType(t *testing.T) {
	if enableManifestStage.GetType() != EnableManifestStageType {
		t.Fatalf("Enable Manifest stage GetType() should be %s, not \"%s\"", EnableManifestStageType, enableManifestStage.GetType())
	}
	if enableManifestStage.Type != EnableManifestStageType {
		t.Fatalf("Enable Manifest stage Type should be %s, not \"%s\"", EnableManifestStageType, enableManifestStage.Type)
	}
}
