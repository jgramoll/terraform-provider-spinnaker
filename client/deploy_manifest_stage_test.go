package client

import (
	"encoding/json"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var deployManifestStage DeployManifestStage

func init() {
	deployManifestStage = *NewDeployManifestStage()
	deployManifestStage.Name = "New Deploy Manifest"
	deployManifestStage.Source = DeployManifestSourceText
	*deployManifestStage.Manifests = append(*deployManifestStage.Manifests, Manifest(deployManifestYaml))
	*deployManifestStage.Manifests = append(*deployManifestStage.Manifests, Manifest(anotherManifestYaml))
}

func TestDeployManifestStageGetType(t *testing.T) {
	if deployManifestStage.GetType() != DeployManifestStageType {
		t.Fatalf("Deploy Manifest stage GetType() should be %s, not \"%s\"", DeployManifestStageType, deployManifestStage.GetType())
	}
	if deployManifestStage.Type != DeployManifestStageType {
		t.Fatalf("Deploy Manifest stage Type should be %s, not \"%s\"", DeployManifestStageType, deployManifestStage.Type)
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
		t.Fatalf("Deploy Manifest not as expected: %s", dmp.DiffPrettyText(diffs))
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
	manifestString := string((*stage.Manifests)[0])
	if manifestString != deployManifestYaml {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(deployManifestYaml, manifestString, true)
		t.Fatalf("manifest not as expected: %s", dmp.DiffPrettyText(diffs))
	}
	manifestString = string((*stage.Manifests)[1])
	if manifestString != anotherManifestYaml {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(anotherManifestYaml, manifestString, true)
		t.Fatalf("manifest not as expected: %s", dmp.DiffPrettyText(diffs))
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

var anotherManifestYaml = `another: 1
`

var deployManifestJson = `{
	"name": "New Deploy Manifest",
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
	"notifications": null,
	"expectedArtifacts": [],
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
		},
		{
			"another": 1
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
