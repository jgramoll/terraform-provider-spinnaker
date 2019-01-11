package client

import (
	"github.com/mitchellh/mapstructure"
)

// BakeStageType bake stage
var BakeStageType StageType = "bake"

func init() {
	stageFactories[BakeStageType] = func(stageMap map[string]interface{}) (Stage, error) {
		stage := NewBakeStage()
		if err := mapstructure.Decode(stageMap, stage); err != nil {
			return nil, err
		}
		return stage, nil
	}
}

// BakeStage for pipeline
type BakeStage struct {
	// TODO why does BaseStage not like mapstructure
	// BaseStage
	Name                 string    `json:"name"`
	RefID                string    `json:"refId"`
	Type                 StageType `json:"type"`
	RequisiteStageRefIds []string  `json:"requisiteStageRefIds"`

	AmiName            string            `json:"amiName"`
	AmiSuffix          string            `json:"amiSuffix,omitempty"`
	BaseAMI            string            `json:"baseAmi,omitempty"`
	BaseLabel          string            `json:"baseLabel"`
	BaseName           string            `json:"baseName,omitempty"`
	BaseOS             string            `json:"baseOs"`
	CloudProviderType  string            `json:"cloudProviderType"`
	ExtendedAttributes map[string]string `json:"extendedAttributes"`
	Rebake             bool              `json:"rebake"`
	Regions            []string          `json:"regions"`
	RequisiteStages    []string          `json:"requisiteStages"`
	StoreType          string            `json:"storeType"`
	TemplateFileName   string            `json:"templateFileName"`
	VarFileName        string            `json:"varFileName,omitempty"`
	VMType             string            `json:"vmType"`
}

// NewBakeStage for pipeline
func NewBakeStage() *BakeStage {
	return &BakeStage{
		// BaseStage: BaseStage{
		Type: BakeStageType,
		// },
	}
}

// GetName for Stage interface
func (s *BakeStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s *BakeStage) GetType() StageType {
	return s.Type
}

// GetRefID for Stage interface
func (s *BakeStage) GetRefID() string {
	return s.RefID
}
