package client

import (
	"testing"
)

var evaluateVariablesStage EvaluateVariablesStage

func init() {
	evaluateVariablesStage = *NewEvaluateVariablesStage()
}

func TestNewEvaluateVariablesStage(t *testing.T) {
	if evaluateVariablesStage.Type != EvaluateVariablesStageType {
		t.Fatalf("Deploy stage type should be %s, not \"%s\"", EvaluateVariablesStageType, evaluateVariablesStage.Type)
	}
}

func TestEvaluateVariablesStageGetName(t *testing.T) {
	name := "New Deploy"
	evaluateVariablesStage.Name = name
	if evaluateVariablesStage.GetName() != name {
		t.Fatalf("Deploy stage GetName() should be %s, not \"%s\"", name, evaluateVariablesStage.GetName())
	}
}

func TestEvaluateVariablesStageGetType(t *testing.T) {
	if evaluateVariablesStage.GetType() != EvaluateVariablesStageType {
		t.Fatalf("Deploy stage GetType() should be %s, not \"%s\"", EvaluateVariablesStageType, evaluateVariablesStage.GetType())
	}
	if evaluateVariablesStage.Type != EvaluateVariablesStageType {
		t.Fatalf("Deploy stage Type should be %s, not \"%s\"", EvaluateVariablesStageType, evaluateVariablesStage.Type)
	}
}
