package client

import (
	"testing"
)

var disableManifestStage DisableManifestStage

func init() {
	disableManifestStage = *NewDisableManifestStage()
}

func TestDisableManifestStageGetName(t *testing.T) {
	name := "New Disable Manifest"
	disableManifestStage.Name = name
	if disableManifestStage.GetName() != name {
		t.Fatalf("Delete Manifest stage GetName() should be %s, not \"%s\"", name, disableManifestStage.GetName())
	}
}

func TestDisableManifestStageGetType(t *testing.T) {
	if disableManifestStage.GetType() != DisableManifestStageType {
		t.Fatalf("Disable Manifest stage GetType() should be %s, not \"%s\"", DisableManifestStageType, disableManifestStage.GetType())
	}
	if disableManifestStage.Type != DisableManifestStageType {
		t.Fatalf("Disable Manifest stage Type should be %s, not \"%s\"", DisableManifestStageType, disableManifestStage.Type)
	}
}
