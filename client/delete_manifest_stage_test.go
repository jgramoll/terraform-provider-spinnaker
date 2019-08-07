package client

import (
	"testing"
)

var deleteManifestStage DeleteManifestStage

func init() {
	deleteManifestStage = *NewDeleteManifestStage()
}

func TestDeleteManifestStageGetName(t *testing.T) {
	name := "New Destroy Server Group"
	deleteManifestStage.Name = name
	if deleteManifestStage.GetName() != name {
		t.Fatalf("Destroy Server Group stage GetName() should be %s, not \"%s\"", name, deleteManifestStage.GetName())
	}
}

func TestDeleteManifestStageGetType(t *testing.T) {
	if deleteManifestStage.GetType() != DeleteManifestStageType {
		t.Fatalf("Destroy Server Group stage GetType() should be %s, not \"%s\"", DeleteManifestStageType, destroyServerGroupStage.GetType())
	}
	if deleteManifestStage.Type != DeleteManifestStageType {
		t.Fatalf("Destroy Server Group stage Type should be %s, not \"%s\"", DeleteManifestStageType, destroyServerGroupStage.Type)
	}
}
