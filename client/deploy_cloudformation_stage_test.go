package client

import (
	"testing"
)

var deployCloudformationStage DeployCloudformationStage

func init() {
	deployCloudformationStage = *NewDeployCloudformationStage()
}

func TestDeployCloudformationStageGetName(t *testing.T) {
	name := "New Deploy Cloudformation"
	deployCloudformationStage.Name = name
	if deployCloudformationStage.GetName() != name {
		t.Fatalf("Deploy Cloudformation stage GetName() should be %s, not \"%s\"", name, deployCloudformationStage.GetName())
	}
}

func TestDeployCloudformationStageGetType(t *testing.T) {
	if deployCloudformationStage.GetType() != DeployCloudformationStageType {
		t.Fatalf("Delete Manifest stage GetType() should be %s, not \"%s\"", DeployCloudformationStageType, deployCloudformationStage.GetType())
	}
	if deployCloudformationStage.Type != DeployCloudformationStageType {
		t.Fatalf("Delete Manifest stage Type should be %s, not \"%s\"", DeployCloudformationStageType, deployCloudformationStage.Type)
	}
}
