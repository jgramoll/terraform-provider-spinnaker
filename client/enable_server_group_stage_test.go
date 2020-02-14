package client

import (
	"testing"
)

var enableServerGroupStage EnableServerGroupStage

func init() {
	enableServerGroupStage = *NewEnableServerGroupStage()
}

func TestEnableServerGroupStageGetName(t *testing.T) {
	name := "New Enable Server Group"
	enableServerGroupStage.Name = name
	if enableServerGroupStage.GetName() != name {
		t.Fatalf("Enable Server Group stage GetName() should be %s, not \"%s\"", name, enableServerGroupStage.GetName())
	}
}

func TestEnableServerGroupStageGetType(t *testing.T) {
	if enableServerGroupStage.GetType() != EnableServerGroupStageType {
		t.Fatalf("Enable Server Group stage GetType() should be %s, not \"%s\"", EnableServerGroupStageType, enableServerGroupStage.GetType())
	}
	if enableServerGroupStage.Type != EnableServerGroupStageType {
		t.Fatalf("Enable Server Group stage Type should be %s, not \"%s\"", EnableServerGroupStageType, enableServerGroupStage.Type)
	}
}
