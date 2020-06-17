package client

import (
	"encoding/json"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var checkPreconditionsStage CheckPreconditionsStage

func init() {
	checkPreconditionsStage = *NewCheckPreconditionsStage()
	checkPreconditionsStage.Name = "Check Preconditions"

	precondition := NewPreconditionStageStatus()
	precondition.Context.StageName = "Manual Judgement"
	precondition.Context.StageStatus = "FAILED_CONTINUE"
	checkPreconditionsStage.Preconditions = append(checkPreconditionsStage.Preconditions, precondition)

	precondition2 := NewPreconditionExpression()
	precondition2.Context.Expression = "this is myexp"
	checkPreconditionsStage.Preconditions = append(checkPreconditionsStage.Preconditions, precondition2)

	precondition3 := NewPreconditionClusterSize()
	precondition3.Context.Credentials = "my-creds"
	precondition3.Context.Expected = 1
	precondition3.Context.Regions = []string{"us-east-2"}
	checkPreconditionsStage.Preconditions = append(checkPreconditionsStage.Preconditions, precondition3)
}

func TestCheckPreconditionsStageGetType(t *testing.T) {
	if checkPreconditionsStage.GetType() != CheckPreconditionsStageType {
		t.Fatalf("Check Preconditions stage GetType() should be %s, not \"%s\"", CheckPreconditionsStageType, checkPreconditionsStage.GetType())
	}
	if checkPreconditionsStage.Type != CheckPreconditionsStageType {
		t.Fatalf("Check Preconditions stage Type should be %s, not \"%s\"", CheckPreconditionsStageType, checkPreconditionsStage.Type)
	}
}

func TestCheckPreconditionsStageSerialize(t *testing.T) {
	b, err := json.MarshalIndent(checkPreconditionsStage, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	result := string(b)
	if result != checkPreconditionsJson {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(checkPreconditionsJson, result, true)
		t.Fatalf("Check Preconditions not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}

func TestCheckPreconditionsStageDeserialize(t *testing.T) {
	var stageMap map[string]interface{}
	err := json.Unmarshal([]byte(checkPreconditionsJson), &stageMap)
	if err != nil {
		t.Fatal(err)
	}
	stageInterface, err := parseCheckPreconditionsStage(stageMap)
	if err != nil {
		t.Fatal(err)
	}
	stage := stageInterface.(*CheckPreconditionsStage)
	if len(stage.Preconditions) != 3 {
		t.Fatalf("Should have preconditions")
	}
}

var checkPreconditionsJson = `{
	"name": "Check Preconditions",
	"refId": "",
	"type": "checkPreconditions",
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
	"preconditions": [
		{
			"failPipeline": true,
			"type": "stageStatus",
			"context": {
				"stageName": "Manual Judgement",
				"stageStatus": "FAILED_CONTINUE"
			}
		},
		{
			"failPipeline": true,
			"type": "expression",
			"context": {
				"expression": "this is myexp"
			}
		},
		{
			"failPipeline": true,
			"type": "clusterSize",
			"context": {
				"credentials": "my-creds",
				"expected": 1,
				"regions": [
					"us-east-2"
				]
			}
		}
	]
}`
