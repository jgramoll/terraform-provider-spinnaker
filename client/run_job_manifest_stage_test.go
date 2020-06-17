package client

import (
	"encoding/json"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var runJobManifestStage RunJobManifestStage

func init() {
	runJobManifestStage = *NewRunJobManifestStage()
	runJobManifestStage.Name = "Run Job (Manifest)"
	runJobManifestStage.Account = "learn-nonprod"
	runJobManifestStage.Application = "bridgelearn"
	runJobManifestStage.CloudProvider = "kubernetes"
	runJobManifestStage.ConsumeArtifactSource = "propertyFile"
	runJobManifestStage.Credentials = "learn-nonprod"
	runJobManifestStage.Manifest = Manifest(runJobManifestYaml)
}

func TestRunJobManifestStageGetType(t *testing.T) {
	if runJobManifestStage.GetType() != RunJobManifestStageType {
		t.Fatalf("Run Job Manifest stage GetType() should be %s, not \"%s\"", RunJobManifestStageType, runJobManifestStage.GetType())
	}
	if runJobManifestStage.Type != RunJobManifestStageType {
		t.Fatalf("Run Job Manifest stage Type should be %s, not \"%s\"", RunJobManifestStageType, runJobManifestStage.Type)
	}
}

func TestRunJobManifestStageSerialize(t *testing.T) {
	b, err := json.MarshalIndent(runJobManifestStage, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	result := string(b)
	if result != runJobManifestJson {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(runJobManifestJson, result, true)
		t.Fatalf("Run Job Manifest not as expected: %s", dmp.DiffPrettyText(diffs))
	}
}

func TestRunJobManifestStageDeserialize(t *testing.T) {
	var stageMap map[string]interface{}
	err := json.Unmarshal([]byte(runJobManifestJson), &stageMap)
	if err != nil {
		t.Fatal(err)
	}
	stageInterface, err := parseRunJobManifestStage(stageMap)
	if err != nil {
		t.Fatal(err)
	}
	stage := stageInterface.(*RunJobManifestStage)
	if string(stage.Manifest) != runJobManifestYaml {
		t.Fatalf("Manifest should be text")
	}
}

var runJobManifestYaml = `apiVersion: batch/v1
kind: Job
metadata:
  name: my-job
  namespace: my-ns
spec:
  template:
    spec:
      containers:
      - command:
        - pwd
        image: echo
        name: halyard
`

var runJobManifestJson = `{
	"name": "Run Job (Manifest)",
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
	"expectedArtifacts": [],
	"account": "learn-nonprod",
	"alias": "runJob",
	"application": "bridgelearn",
	"cloudProvider": "kubernetes",
	"consumeArtifactSource": "propertyFile",
	"credentials": "learn-nonprod",
	"manifest": {
		"apiVersion": "batch/v1",
		"kind": "Job",
		"metadata": {
			"name": "my-job",
			"namespace": "my-ns"
		},
		"spec": {
			"template": {
				"spec": {
					"containers": [
						{
							"command": [
								"pwd"
							],
							"image": "echo",
							"name": "halyard"
						}
					]
				}
			}
		}
	},
	"propertyFile": "",
	"source": ""
}`
