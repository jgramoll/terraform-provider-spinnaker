package client

import (
	"testing"
)

var undoRolloutManifestStage UndoRolloutManifestStage

func init() {
	undoRolloutManifestStage = *NewUndoRolloutManifestStage()
}

func TestUndoRolloutManifestStageGetName(t *testing.T) {
	name := "New Undo Rollout Manifest"
	undoRolloutManifestStage.Name = name
	if undoRolloutManifestStage.GetName() != name {
		t.Fatalf("Undo Rollout Manifest stage GetName() should be %s, not \"%s\"", name, undoRolloutManifestStage.GetName())
	}
}

func TestUndoRolloutManifestStageGetType(t *testing.T) {
	if undoRolloutManifestStage.GetType() != UndoRolloutManifestStageType {
		t.Fatalf("Undo Rollout Manifest stage GetType() should be %s, not \"%s\"", UndoRolloutManifestStageType, undoRolloutManifestStage.GetType())
	}
	if undoRolloutManifestStage.Type != UndoRolloutManifestStageType {
		t.Fatalf("Undo Rollout Manifest stage Type should be %s, not \"%s\"", UndoRolloutManifestStageType, undoRolloutManifestStage.Type)
	}
}
