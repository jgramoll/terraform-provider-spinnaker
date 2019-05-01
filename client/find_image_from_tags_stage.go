package client

import (
	"github.com/mitchellh/mapstructure"
)

// FindImageFromTagsStageType bake stage
var FindImageFromTagsStageType StageType = "findImageFromTags"

func init() {
	stageFactories[FindImageFromTagsStageType] = parseFindImageStage
}

// FindImageFromTagsStage for pipeline
type FindImageFromTagsStage struct {
	BaseStage `mapstructure:",squash"`

	CloudProvider     string            `json:"cloudProvider"`
	CloudProviderType string            `json:"cloudProviderType"`
	PackageName       string            `json:"packageName"`
	Regions           []string          `json:"regions"`
	Tags              map[string]string `json:"tags"`
}

// FindImageFromTagsStage for pipeline
type FindImageFromTagsStage struct {
	*serializableFindImageFromTagsStage
	Notifications *[]*Notification `json:"notifications"`
}

func newserializableFindImageStage() *serializableFindImageFromTagsStage {
	return &serializableFindImageFromTagsStage{
		Type:                 FindImageFromTagsStageType,
		FailPipeline:         true,
		RequisiteStageRefIds: []string{},
	}
}

// NewFindImageStage for pipeline
func NewFindImageStage() *FindImageFromTagsStage {
	return &FindImageFromTagsStage{
		BaseStage: *newBaseStage(FindImageFromTagsStageType),
	}
}

func parseFindImageStage(stageMap map[string]interface{}) (Stage, error) {
	stage := NewFindImageStage()
	if err := stage.parseBaseStage(stageMap); err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(stageMap, stage); err != nil {
		return nil, err
	}
	return stage, nil
}
