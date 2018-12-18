package client

// BakeStageType bake stage
var BakeStageType StageType = "bake"

func init() {
	stageFactories[BakeStageType] = func() interface{} {
		return NewBakeStage()
	}
}

// BakeStage for pipeline
type BakeStage struct {
	// TODO why does BaseStage not like mapstructure
	// BaseStage
	Name  string    `json:"name"`
	RefID string    `json:"refId"`
	Type  StageType `json:"type"`

	AmiName            string            `json:"amiName"`
	AmiSuffix          string            `json:"amiSuffix"`
	BaseAMI            string            `json:"baseAmi"`
	BaseLabel          string            `json:"baseLabel"`
	BaseName           string            `json:"baseName"`
	BaseOS             string            `json:"baseOs"`
	CloudProviderType  string            `json:"cloudProviderType"`
	ExtendedAttributes map[string]string `json:"extendedAttributes"`
	Regions            []string          `json:"regions"`
	RequisiteStages    []string          `json:"requisiteStages"`
	StoreType          string            `json:"storeType"`
	TemplateFileName   string            `json:"templateFileName"`
	VarFileName        string            `json:"varFileName"`
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
func (s BakeStage) GetName() string {
	return s.Name
}

// GetType for Stage interface
func (s BakeStage) GetType() StageType {
	return s.Type
}
