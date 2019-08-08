package client

import (
	"encoding/json"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var deployManifestStage DeployManifestStage

func init() {
	deployManifestStage = *NewDeployManifestStage()
	deployManifestStage.Source = DeployManifestSourceText
	*deployManifestStage.Manifests = append(*deployManifestStage.Manifests, deployManifestYaml)
}

func TestDeployManifestStageGetName(t *testing.T) {
	name := "New Deploy Manifest"
	deployManifestStage.Name = name
	if deployManifestStage.GetName() != name {
		t.Fatalf("Destroy Deploy Manifest stage GetName() should be %s, not \"%s\"", name, deployManifestStage.GetName())
	}
}

func TestDeployManifestStageGetType(t *testing.T) {
	if deployManifestStage.GetType() != DeployManifestStageType {
		t.Fatalf("Destroy Deploy Manifest stage GetType() should be %s, not \"%s\"", DeployManifestStageType, deployManifestStage.GetType())
	}
	if deployManifestStage.Type != DeployManifestStageType {
		t.Fatalf("Destroy Deploy Manifest stage Type should be %s, not \"%s\"", DeployManifestStageType, deployManifestStage.Type)
	}
}

func TestDeployManifestStageSerialize(t *testing.T) {
	b, err := json.MarshalIndent(deployManifestStage, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	result := string(b)
	if result != deployManifestJson {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(deployManifestJson, result, true)
		t.Fatalf("job definition not expected: %s", dmp.DiffPrettyText(diffs))
	}
}

func TestDeployManifestStageDeserialize(t *testing.T) {
	var stageMap map[string]interface{}
	err := json.Unmarshal([]byte(deployManifestJson), &stageMap)
	if err != nil {
		t.Fatal(err)
	}
	stageInterface, err := parseDeployManifestStage(stageMap)
	if err != nil {
		t.Fatal(err)
	}
	stage := stageInterface.(*DeployManifestStage)
	if stage.Source != DeployManifestSourceText {
		t.Fatalf("Source should be text")
	}
	if len(*stage.Manifests) == 0 {
		t.Fatalf("Should have manifest")
	}
	manifestString := (*stage.Manifests)[0]
	if manifestString != deployManifestYaml {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(deployManifestYaml, manifestString, true)
		t.Fatalf("job definition not expected: %s", dmp.DiffPrettyText(diffs))
	}
}

var deployManifestYaml = `apiVersion: batch/v1
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

var deployManifestJson = `{
	"name": "",
	"refId": "",
	"type": "deployManifest",
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
	"account": "",
	"cloudProvider": "",
	"manifestArtifactAccount": "",
	"manifests": [
		{
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
		}
	],
	"moniker": null,
	"relationships": {
		"loadBalancers": [],
		"securityGroups": []
	},
	"skipExpressionEvaluation": false,
	"source": "text",
	"trafficManagement": {
		"enabled": false,
		"options": {
			"enableTraffic": false
		}
	}
}`
