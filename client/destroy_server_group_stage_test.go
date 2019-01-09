package client

import (
	"testing"
)

var destroyServerGroupStage DestroyServerGroupStage

func init() {
	destroyServerGroupStage = *NewDestroyServerGroupStage()
}

func TestDestroyServerGroupStageGetName(t *testing.T) {
	name := "New Destroy Server Group"
	destroyServerGroupStage.Name = name
	if destroyServerGroupStage.GetName() != name {
		t.Fatalf("Destroy Server Group stage GetName() should be %s, not \"%s\"", name, destroyServerGroupStage.GetName())
	}
}

func TestDestroyServerGroupStageGetType(t *testing.T) {
	if destroyServerGroupStage.GetType() != DestroyServerGroupType {
		t.Fatalf("Destroy Server Group stage GetType() should be %s, not \"%s\"", DestroyServerGroupType, destroyServerGroupStage.GetType())
	}
	if destroyServerGroupStage.Type != DestroyServerGroupType {
		t.Fatalf("Destroy Server Group stage Type should be %s, not \"%s\"", DestroyServerGroupType, destroyServerGroupStage.Type)
	}
}
