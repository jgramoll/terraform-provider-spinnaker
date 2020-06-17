package client

import (
	"testing"
)

var bakeManifestStage BakeManifestStage

func init() {
	bakeManifestStage = *NewBakeManifestStage()
	bakeManifestStage.Name = "New Deploy Manifest"
}

func TestBakeManifestStageGetType(t *testing.T) {
	if bakeManifestStage.GetType() != BakeManifestStageType {
		t.Fatalf("Deploy Manifest stage GetType() should be %s, not \"%s\"", BakeManifestStageType, bakeManifestStage.GetType())
	}
	if bakeManifestStage.Type != BakeManifestStageType {
		t.Fatalf("Deploy Manifest stage Type should be %s, not \"%s\"", BakeManifestStageType, bakeManifestStage.Type)
	}
}
