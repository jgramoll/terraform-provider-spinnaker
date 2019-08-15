package client

import (
	"testing"
)

var manualJudgment ManualJudgmentStage

func init() {
	manualJudgment = *NewManualJudgmentStage()
}

func TestNewManualJudgmentStage(t *testing.T) {
	if manualJudgment.Type != ManualJudgmentStageType {
		t.Fatalf("ManualJudgment stage type should be %s, not \"%s\"", ManualJudgmentStageType, manualJudgment.Type)
	}
}

func TestManualJudgmentStageGetName(t *testing.T) {
	name := "New Manual Judgment"
	manualJudgment.Name = name
	if manualJudgment.GetName() != name {
		t.Fatalf("Manual Judgment stage GetName() should be %s, not \"%s\"", name, manualJudgment.GetName())
	}
}

func TestManualJudgmentStageGetType(t *testing.T) {
	if manualJudgment.GetType() != ManualJudgmentStageType {
		t.Fatalf("Manual Judgment stage GetType() should be %s, not \"%s\"", ManualJudgmentStageType, manualJudgment.GetType())
	}
	if manualJudgment.Type != ManualJudgmentStageType {
		t.Fatalf("Manual Judgment stage Type should be %s, not \"%s\"", ManualJudgmentStageType, manualJudgment.Type)
	}
}
