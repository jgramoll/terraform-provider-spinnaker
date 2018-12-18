package client

import (
	"testing"
)

var bakeStage BakeStage

func init() {
	bakeStage = *NewBakeStage()
}

func TestNewBakeStage(t *testing.T) {
	if bakeStage.Type != BakeStageType {
		t.Fatalf("Bake stage type should be %s, not \"%s\"", BakeStageType, bakeStage.Type)
	}
}

func TestBakeStageGetName(t *testing.T) {
	name := "New Bake"
	bakeStage.Name = name
	if bakeStage.GetName() != name {
		t.Fatalf("Bake stage GetName() should be %s, not \"%s\"", name, bakeStage.GetName())
	}
	if bakeStage.GetName() != bakeStage.Name {
		t.Fatalf("Bake stage GetName() should be %s, not \"%s\"", bakeStage.Name, bakeStage.GetName())
	}
}

func TestBakeStageGetType(t *testing.T) {
	if bakeStage.GetType() != BakeStageType {
		t.Fatalf("Bake stage GetType() should be %s, not \"%s\"", BakeStageType, bakeStage.GetType())
	}
	if bakeStage.GetType() != bakeStage.Type {
		t.Fatalf("Bake stage GetType() should be %s, not \"%s\"", bakeStage.Type, bakeStage.GetType())
	}
}
