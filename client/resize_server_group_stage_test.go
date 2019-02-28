package client

import (
	"testing"
)

var resizeServerGroupStage ResizeServerGroupStage

func init() {
	resizeServerGroupStage = *NewResizeServerGroupStage()
}

func TestResizeServerGroupStageGetName(t *testing.T) {
	name := "New Resize"
	resizeServerGroupStage.Name = name
	if resizeServerGroupStage.GetName() != name {
		t.Fatalf("Resize Server Group stage GetName() should be %s, not \"%s\"", name, resizeServerGroupStage.GetName())
	}
}

func TestResizeServerGroupStageGetType(t *testing.T) {
	if resizeServerGroupStage.GetType() != ResizeServerGroupStageType {
		t.Fatalf("Resize Server Group stage GetType() should be %s, not \"%s\"", ResizeServerGroupStageType, resizeServerGroupStage.GetType())
	}
	if resizeServerGroupStage.Type != ResizeServerGroupStageType {
		t.Fatalf("Resize Server Group stage Type should be %s, not \"%s\"", ResizeServerGroupStageType, resizeServerGroupStage.Type)
	}
}
