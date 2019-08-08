package client

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

// DeployManifestStageType deploy manifest stage
var DeployManifestStageType StageType = "deployManifest"

func init() {
	stageFactories[DeployManifestStageType] = parseDeployManifestStage
}

type DeployManifestStage struct {
	// BaseStage
	Name                              string                `json:"name"`
	RefID                             string                `json:"refId"`
	Type                              StageType             `json:"type"`
	RequisiteStageRefIds              []string              `json:"requisiteStageRefIds"`
	SendNotifications                 bool                  `json:"sendNotifications"`
	StageEnabled                      *StageEnabled         `json:"stageEnabled"`
	CompleteOtherBranchesThenFail     bool                  `json:"completeOtherBranchesThenFail"`
	ContinuePipeline                  bool                  `json:"continuePipeline"`
	FailOnFailedExpressions           bool                  `json:"failOnFailedExpressions"`
	FailPipeline                      bool                  `json:"failPipeline"`
	OverrideTimeout                   bool                  `json:"overrideTimeout"`
	RestrictExecutionDuringTimeWindow bool                  `json:"restrictExecutionDuringTimeWindow"`
	RestrictedExecutionWindow         *StageExecutionWindow `json:"restrictedExecutionWindow"`
	// End BaseStage

	Account                  string               `json:"account"`
	CloudProvider            string               `json:"cloudProvider"`
	ManifestArtifactAccount  string               `json:"manifestArtifactAccount"`
	Manifests                *DeployManifests     `json:"manifests"`
	Moniker                  *Moniker             `json:"moniker"`
	Relationships            *Relationships       `json:"relationships"`
	SkipExpressionEvaluation bool                 `json:"skipExpressionEvaluation"`
	Source                   DeployManifestSource `json:"source"` // TODO enum
	TrafficManagement        *TrafficManagement   `json:"trafficManagement"`
}

func NewDeployManifestStage() *DeployManifestStage {
	return &DeployManifestStage{
		Type:                 DeployManifestStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
		Manifests:            NewDeployManifests(),
		Relationships:        NewRelationships(),
		TrafficManagement:    NewTrafficManagement(),
	}
}

// GetName for Stage interface
func (s *DeployManifestStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *DeployManifestStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *DeployManifestStage) GetRefID() string {
	return s.RefID
}

func parseDeployManifestStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewDeployManifestStage()
	manifestInterface, ok := stageMap["manifests"].([]interface{})
	if !ok {
		println(reflect.TypeOf(stageMap["manifests"]).String())
		return nil, fmt.Errorf("Could not parse deploy manifest manifests: %v", stageMap["manifests"])
	}
	sourceString, ok := stageMap["source"].(string)
	if !ok {
		return nil, fmt.Errorf("Could not parse deploy manifest source: %v", stageMap["source"])
	}

	delete(stageMap, "manifests")
	delete(stageMap, "source")
	if err := mapstructure.Decode(stageMap, stage); err != nil {
		println("ASJdf")
		return nil, err
	}

	source, err := ParseDeployManifestSource(sourceString)
	if err != nil {
		return nil, err
	}
	stage.Source = source

	manifests, err := ParseDeployManifests(manifestInterface)
	if err != nil {
		return nil, err
	}
	stage.Manifests = manifests

	return stage, nil
}
