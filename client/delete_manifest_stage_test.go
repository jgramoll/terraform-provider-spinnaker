package client

import (
	"testing"
)

var deleteManifestStage DeleteManifestStage

func init() {
	deleteManifestStage = *NewDeleteManifestStage()
}

func TestDeleteManifestStageGetName(t *testing.T) {
	name := "New Destroy Delete Manifest"
	deleteManifestStage.Name = name
	if deleteManifestStage.GetName() != name {
		t.Fatalf("Destroy Delete Manifest stage GetName() should be %s, not \"%s\"", name, deleteManifestStage.GetName())
	}
}

func TestDeleteManifestStageGetType(t *testing.T) {
	if deleteManifestStage.GetType() != DeleteManifestStageType {
		t.Fatalf("Destroy Delete Manifest stage GetType() should be %s, not \"%s\"", DeleteManifestStageType, deleteManifestStage.GetType())
	}
	if deleteManifestStage.Type != DeleteManifestStageType {
		t.Fatalf("Destroy Delete Manifest stage Type should be %s, not \"%s\"", DeleteManifestStageType, deleteManifestStage.Type)
	}
}
