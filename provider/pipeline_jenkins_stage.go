package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type jenkinsStage struct {
	Name                 string           `mapstructure:"name"`
	RefID                string           `mapstructure:"ref_id"`
	Type                 client.StageType `mapstructure:"type"`
	RequisiteStageRefIds []string         `mapstructure:"requisite_stage_ref_ids"`

	CompleteOtherBranchesThenFail bool              `mapstructure:"complete_other_branches_then_fail"`
	ContinuePipeline              bool              `mapstructure:"continue_pipeline"`
	FailPipeline                  bool              `mapstructure:"fail_pipeline"`
	Job                           string            `mapstructure:"job"`
	MarkUnstableAsSuccessful      bool              `mapstructure:"mark_unstable_as_successful"`
	Master                        string            `mapstructure:"master"`
	Parameters                    map[string]string `mapstructure:"parameters"`
	PropertyFile                  string            `mapstructure:"property_file"`
}

func newJenkinsStage() *jenkinsStage {
	return &jenkinsStage{Type: client.JenkinsStageType}
}

func (s *jenkinsStage) toClientStage() client.Stage {
	cs := client.JenkinsStage(*s)
	return &cs
}

// TODO can we just update the ptr?
func (s *jenkinsStage) fromClientStage(cs client.Stage) stage {
	newStage := jenkinsStage(*(cs.(*client.JenkinsStage)))
	// s = &newStage
	return stage(&newStage)
}

func (s *jenkinsStage) SetResourceData(d *schema.ResourceData) {
	// TODO
	d.Set("name", s.Name)
}

func (s *jenkinsStage) SetRefID(id string) {
	s.RefID = id
}

func (s *jenkinsStage) GetRefID() string {
	return s.RefID
}
