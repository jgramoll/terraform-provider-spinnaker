package client

import (
	"testing"
)

var rollbackClusterStage RollbackClusterStage

func init() {
	rollbackClusterStage = *NewRollbackClusterStage()
}

func TestRollbackClusterStageGetName(t *testing.T) {
	name := "New Rollback"
	rollbackClusterStage.Name = name
	if rollbackClusterStage.GetName() != name {
		t.Fatalf("Rollback Cluster stage GetName() should be %s, not \"%s\"", name, rollbackClusterStage.GetName())
	}
}

func TestRollbackClusterStageGetType(t *testing.T) {
	if rollbackClusterStage.GetType() != RollbackClusterStageType {
		t.Fatalf("Rollback Cluster stage GetType() should be %s, not \"%s\"", RollbackClusterStageType, rollbackClusterStage.GetType())
	}
	if rollbackClusterStage.Type != RollbackClusterStageType {
		t.Fatalf("Rollback Cluster stage Type should be %s, not \"%s\"", RollbackClusterStageType, rollbackClusterStage.Type)
	}
}
