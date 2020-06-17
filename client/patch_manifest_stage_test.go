package client

import (
	"testing"
)

var patchManifestStage PatchManifestStage

func init() {
	patchManifestStage = *NewPatchManifestStage()
	patchManifestStage.Name = "New Patch Manifest"
}

func TestPatchManifestStageGetType(t *testing.T) {
	if patchManifestStage.GetType() != PatchManifestStageType {
		t.Fatalf("Patch Manifest stage GetType() should be %s, not \"%s\"", PatchManifestStageType, patchManifestStage.GetType())
	}
	if patchManifestStage.Type != PatchManifestStageType {
		t.Fatalf("Patch Manifest stage Type should be %s, not \"%s\"", PatchManifestStageType, patchManifestStage.Type)
	}
}
