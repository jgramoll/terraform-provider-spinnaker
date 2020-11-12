package client

import (
	"encoding/json"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var deployCloudformationStage DeployCloudformationStage

func init() {
	deployCloudformationStage = *NewDeployCloudformationStage()
	deployCloudformationStage.Name = "New Deploy Cloudformation"
}

func TestDeployCloudformationStageGetName(t *testing.T) {
	name := "New Deploy Cloudformation"
	if deployCloudformationStage.GetName() != name {
		t.Fatalf("Deploy Cloudformation stage GetName() should be %s, not \"%s\"", name, deployCloudformationStage.GetName())
	}
}

func TestDeployCloudformationStageGetType(t *testing.T) {
	if deployCloudformationStage.GetType() != DeployCloudformationStageType {
		t.Fatalf("Deploy Cloudformation stage GetType() should be %s, not \"%s\"", DeployCloudformationStageType, deployCloudformationStage.GetType())
	}
	if deployCloudformationStage.Type != DeployCloudformationStageType {
		t.Fatalf("Deploy Cloudformation stage Type should be %s, not \"%s\"", DeployCloudformationStageType, deployCloudformationStage.Type)
	}
}

func TestDeployCloudformationStageSerialize(t *testing.T) {
	b, err := json.MarshalIndent(deployCloudformationStage, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	result := string(b)
	if result != deployCloudformationJSON {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(deployCloudformationJSON, result, true)
		t.Fatalf("Deploy Cloudformation not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}

func TestDeployCloudformationStageDeserialize(t *testing.T) {
	var stageMap map[string]interface{}
	err := json.Unmarshal([]byte(deployCloudformationJSON), &stageMap)
	if err != nil {
		t.Fatal(err)
	}
	stageInterface, err := parseDeployCloudformationStage(stageMap)
	if err != nil {
		t.Fatal(err)
	}
	stage := stageInterface.(*DeployCloudformationStage)
	if stage.Source != DeployCloudformationSourceText {
		t.Fatalf("Should have source type text")
	}
}

var deployCloudformationJSON = `{
	"name": "New Deploy Cloudformation",
	"refId": "",
	"type": "deployCloudFormation",
	"requisiteStageRefIds": [],
	"sendNotifications": false,
	"stageEnabled": null,
	"completeOtherBranchesThenFail": false,
	"continuePipeline": false,
	"failOnFailedExpressions": false,
	"failPipeline": true,
	"overrideTimeout": false,
	"restrictExecutionDuringTimeWindow": false,
	"restrictedExecutionWindow": null,
	"notifications": null,
	"expectedArtifacts": [],
	"requiredArtifactIds": [],
	"actionOnReplacement": "",
	"capabilities": null,
	"changeSetName": "",
	"credentials": "",
	"executeChangeSet": false,
	"isChangeSet": false,
	"parameters": null,
	"regions": null,
	"roleARN": "",
	"source": "text",
	"stackArtifact": null,
	"stackName": "",
	"tags": null,
	"templateBody": null
}`
