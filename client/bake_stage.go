package client

import (
	"github.com/mitchellh/mapstructure"
)

// BakeStageType bake stage
var BakeStageType StageType = "bake"

func init() {
	stageFactories[BakeStageType] = parseBakeStage
}

// BakeStage for pipeline
type BakeStage struct {
	BaseStage `mapstructure:",squash"`

	AmiName            string            `json:"amiName"`
	AmiSuffix          string            `json:"amiSuffix,omitempty"`
	BaseAMI            string            `json:"baseAmi,omitempty"`
	BaseLabel          string            `json:"baseLabel"`
	BaseName           string            `json:"baseName,omitempty"`
	BaseOS             string            `json:"baseOs"`
	CloudProviderType  string            `json:"cloudProviderType"`
	ExtendedAttributes map[string]string `json:"extendedAttributes"`
	Package            string            `json:"package"`
	Rebake             bool              `json:"rebake"`
	Region             string            `json:"region"`
	Regions            []string          `json:"regions,omitempty"`
	StoreType          string            `json:"storeType"`
	TemplateFileName   string            `json:"templateFileName"`
	User               string            `json:"user,omitempty"`
	VarFileName        string            `json:"varFileName,omitempty"`
	VMType             string            `json:"vmType"`
}

// NewBakeStage for pipeline
func NewBakeStage() *BakeStage {
	return &BakeStage{
		BaseStage: *newBaseStage(BakeStageType),
	}
}

func parseBakeStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewBakeStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
