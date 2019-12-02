package client

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// DeployManifestStageType deploy manifest stage
var DeployManifestStageType StageType = "deployManifest"

func init() {
	stageFactories[DeployManifestStageType] = parseDeployManifestStage
}

type DeployManifestStage struct {
	BaseStage `mapstructure:",squash"`

	Account                  string               `json:"account"`
	NamespaceOverride        string               `json:"namespaceOverride,omitempty"`
	CloudProvider            string               `json:"cloudProvider"`
	ManifestArtifactAccount  string               `json:"manifestArtifactAccount"`
	Manifests                *Manifests           `json:"manifests"`
	Moniker                  *Moniker             `json:"moniker"`
	Relationships            *Relationships       `json:"relationships"`
	SkipExpressionEvaluation bool                 `json:"skipExpressionEvaluation"`
	Source                   DeployManifestSource `json:"source"`
	TrafficManagement        *TrafficManagement   `json:"trafficManagement"`
}

func NewDeployManifestStage() *DeployManifestStage {
	return &DeployManifestStage{
		BaseStage: *newBaseStage(DeployManifestStageType),

		Manifests:         NewManifests(),
		Relationships:     NewRelationships(),
		TrafficManagement: NewTrafficManagement(),
	}
}

func parseDeployManifestStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewDeployManifestStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	manifestInterface, ok := stageMap["manifests"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("Could not parse deploy manifest manifests: %v", stageMap["manifests"])
	}
	manifests, err := ParseDeployManifests(manifestInterface)
	if err != nil {
		return nil, err
	}
	stage.Manifests = manifests
	delete(stageMap, "manifests")

	sourceString, ok := stageMap["source"].(string)
	if !ok {
		return nil, fmt.Errorf("Could not parse deploy manifest source: %v", stageMap["source"])
	}
	source, err := ParseDeployManifestSource(sourceString)
	if err != nil {
		return nil, err
	}
	stage.Source = source
	delete(stageMap, "source")

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
