package client

import (
	"encoding/json"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var expectedArtifactsStage RunJobManifestStage

func init() {
	expectedArtifact := NewManifestExpectedArtifact()
	expectedArtifact.DefaultArtifact.CustomKind = true
	expectedArtifact.DefaultArtifact.ID = "ck"
	expectedArtifact.DisplayName = "dname"
	expectedArtifact.ID = "aid"
	expectedArtifact.MatchArtifact.ID = "mad"
	expectedArtifact.MatchArtifact.Location = "mloc"
	expectedArtifact.MatchArtifact.Name = "mname"
	expectedArtifact.MatchArtifact.Reference = "mref"
	expectedArtifact.MatchArtifact.Type = "mtype"

	expectedArtifactsStage = *NewRunJobManifestStage()
	expectedArtifactsStage.Name = "Run test"
	expectedArtifactsStage.Manifest = "foo"
	*expectedArtifactsStage.ExpectedArtifacts = append(*expectedArtifactsStage.ExpectedArtifacts, expectedArtifact)
}

func TestExpectedArtifactsSerialize(t *testing.T) {
	b, err := json.MarshalIndent(expectedArtifactsStage, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	result := string(b)
	if result != expectedArtifactsJSON {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expectedArtifactsJSON, result, true)
		t.Fatalf("Expected artifacts not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}

func TestExpectedArtifactsDesrialize(t *testing.T) {
	var stageMap map[string]interface{}
	err := json.Unmarshal([]byte(expectedArtifactsJSON), &stageMap)
	if err != nil {
		t.Fatal(err)
	}
	stageInterface, err := parseRunJobManifestStage(stageMap)
	if err != nil {
		t.Fatal(err)
	}
	stage := stageInterface.(*RunJobManifestStage)
	if len(*stage.ExpectedArtifacts) != 1 {
		t.Fatalf("Should have expected artifacts")
	}
}

var expectedArtifactsJSON = `{
	"name": "Run test",
	"refId": "",
	"type": "runJobManifest",
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
	"expectedArtifacts": [
		{
			"defaultArtifact": {
				"customKind": true,
				"id": "ck"
			},
			"displayName": "dname",
			"id": "aid",
			"matchArtifact": {
				"customKind": false,
				"id": "mad",
				"location": "mloc",
				"name": "mname",
				"reference": "mref",
				"type": "mtype"
			},
			"useDefaultArtifact": false,
			"usePriorArtifact": false
		}
	],
	"requiredArtifactIds": [],
	"account": "",
	"alias": "runJob",
	"application": "",
	"cloudProvider": "",
	"consumeArtifactSource": "",
	"credentials": "",
	"manifest": "foo",
	"propertyFile": "",
	"source": ""
}`
