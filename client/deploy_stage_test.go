package client

import (
	"testing"
)

var deployStage DeployStage

func init() {
	deployStage = *NewDeployStage()
}

func TestNewDeployStage(t *testing.T) {
	if deployStage.Type != DeployStageType {
		t.Fatalf("Deploy stage type should be %s, not \"%s\"", DeployStageType, deployStage.Type)
	}
}

func TestDeployStageGetName(t *testing.T) {
	name := "New Deploy"
	deployStage.Name = name
	if deployStage.GetName() != name {
		t.Fatalf("Deploy stage GetName() should be %s, not \"%s\"", name, deployStage.GetName())
	}
}

func TestDeployStageGetType(t *testing.T) {
	if deployStage.GetType() != DeployStageType {
		t.Fatalf("Deploy stage GetType() should be %s, not \"%s\"", DeployStageType, deployStage.GetType())
	}
	if deployStage.Type != DeployStageType {
		t.Fatalf("Deploy stage Type should be %s, not \"%s\"", DeployStageType, deployStage.Type)
	}
}
