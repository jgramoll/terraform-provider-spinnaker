package client

import (
	"testing"
)

var findImageStage FindImageFromTagsStage

func init() {
	findImageStage = *NewFindImageStage()
}

func TestNewFindImageStage(t *testing.T) {
	if findImageStage.Type != FindImageFromTagsStageType {
		t.Fatalf("Deploy stage type should be %s, not \"%s\"", FindImageFromTagsStageType, findImageStage.Type)
	}
}

func TestFindImageStageGetName(t *testing.T) {
	name := "New Deploy"
	findImageStage.Name = name
	if findImageStage.GetName() != name {
		t.Fatalf("Deploy stage GetName() should be %s, not \"%s\"", name, findImageStage.GetName())
	}
}

func TestFindImageStageGetType(t *testing.T) {
	if findImageStage.GetType() != FindImageFromTagsStageType {
		t.Fatalf("Deploy stage GetType() should be %s, not \"%s\"", FindImageFromTagsStageType, findImageStage.GetType())
	}
	if findImageStage.Type != FindImageFromTagsStageType {
		t.Fatalf("Deploy stage Type should be %s, not \"%s\"", FindImageFromTagsStageType, findImageStage.Type)
	}
}
