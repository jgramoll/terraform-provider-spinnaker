package client

import (
	"testing"
)

var disableServerGroupStage DisableServerGroupStage

func init() {
	disableServerGroupStage = *NewDisableServerGroupStage()
}

func TestDisableServerGroupStageGetName(t *testing.T) {
	name := "New Disable Server Group"
	disableServerGroupStage.Name = name
	if disableServerGroupStage.GetName() != name {
		t.Fatalf("Disable Server Group stage GetName() should be %s, not \"%s\"", name, disableServerGroupStage.GetName())
	}
}

func TestDisableServerGroupStageGetType(t *testing.T) {
	if disableServerGroupStage.GetType() != DisableServerGroupStageType {
		t.Fatalf("Disable Server Group stage GetType() should be %s, not \"%s\"", DisableServerGroupStageType, disableServerGroupStage.GetType())
	}
	if disableServerGroupStage.Type != DisableServerGroupStageType {
		t.Fatalf("Disable Server Group stage Type should be %s, not \"%s\"", DisableServerGroupStageType, disableServerGroupStage.Type)
	}
}
