package client

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// DeployCloudformationStageType deploy manifest stage
var DeployCloudformationStageType StageType = "deployCloudFormation"

func init() {
	stageFactories[DeployCloudformationStageType] = parseDeployCloudformationStage
}

// DeployCloudformationStage deploy cloudforamtion stage
type DeployCloudformationStage struct {
	BaseStage `mapstructure:",squash"`

	ActionOnReplacement string                     `json:"actionOnReplacement"`
	Capabilities        []string                   `json:"capabilities"`
	ChangeSetName       string                     `json:"changeSetName"`
	Credentials         string                     `json:"credentials"`
	ExecuteChangeSet    bool                       `json:"executeChangeSet"`
	IsChangeSet         bool                       `json:"isChangeSet"`
	Parameters          map[string]string          `json:"parameters"`
	Regions             []string                   `json:"regions"`
	RoleARN             string                     `json:"roleARN"`
	Source              DeployCloudformationSource `json:"source"`
	StackArtifact       *StackArtifact             `json:"stackArtifact"`
	StackName           string                     `json:"stackName"`
	Tags                map[string]string          `json:"tags"`
	TemplateBody        []string                   `json:"templateBody"`
}

// NewDeployCloudformationStage deploy cloudformation stage
func NewDeployCloudformationStage() *DeployCloudformationStage {
	return &DeployCloudformationStage{
		BaseStage: *newBaseStage(DeployCloudformationStageType),
		Source:    DeployCloudformationSourceText,
	}
}

func parseDeployCloudformationStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewDeployCloudformationStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	sourceString, ok := stageMap["source"].(string)
	if !ok {
		return nil, fmt.Errorf("Could not parse cloudformation source %v", stageMap["source"])
	}
	source, err := ParseDeployCloudformationSource(sourceString)
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
