package client

import (
	"github.com/mitchellh/mapstructure"
)

// BakeStageType bake stage
var BakeStageType StageType = "bake"

func init() {
	stageFactories[BakeStageType] = parseBakeStage
}

type serializableBakeStage struct {
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

// BakeStage for pipeline
type BakeStage struct {
	*serializableBakeStage
	Notifications *[]*Notification `json:"notifications"`
}

func newSerializableBakeStage() *serializableBakeStage {
	return &serializableBakeStage{
		Type:                 BakeStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
	}
}

// NewBakeStage for pipeline
func NewBakeStage() *BakeStage {
	return &BakeStage{
		serializableBakeStage: newSerializableBakeStage(),
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

func parseBakeStage(stageMap map[string]interface{}) (Stage, error) {
	stage := newSerializableBakeStage()
	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}

	notifications, err := parseNotifications(stageMap["notifications"])
	if err != nil {
		return nil, err
	}
	return &BakeStage{
		serializableBakeStage: stage,
		Notifications:         notifications,
	}, nil
}
