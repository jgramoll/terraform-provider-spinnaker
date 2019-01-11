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
	Notifications        []*notification  `mapstructure:"notification"`

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
	cs := client.NewJenkinsStage()
	cs.Name = s.Name
	cs.RefID = s.RefID
	cs.RequisiteStageRefIds = s.RequisiteStageRefIds
	cs.SendNotifications = len(s.Notifications) > 0

	for _, n := range s.Notifications {
		cs.Notifications = append(cs.Notifications, n.toClientNotification(client.NotificationLevelStage))
	}

	cs.CompleteOtherBranchesThenFail = s.CompleteOtherBranchesThenFail
	cs.ContinuePipeline = s.ContinuePipeline
	cs.FailPipeline = s.FailPipeline
	cs.Job = s.Job
	cs.MarkUnstableAsSuccessful = s.MarkUnstableAsSuccessful
	cs.Master = s.Master
	cs.Parameters = s.Parameters
	cs.PropertyFile = s.PropertyFile

	return cs
}

// TODO can we just update the ptr?
func (s *jenkinsStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.JenkinsStage)
	newStage := newJenkinsStage()
	newStage.Name = clientStage.Name
	newStage.RefID = clientStage.RefID
	newStage.RequisiteStageRefIds = clientStage.RequisiteStageRefIds

	for _, cn := range clientStage.Notifications {
		newStage.Notifications = append(newStage.Notifications, notification{}.fromClientNotification(cn))
	}

	newStage.CompleteOtherBranchesThenFail = clientStage.CompleteOtherBranchesThenFail
	newStage.ContinuePipeline = clientStage.ContinuePipeline
	newStage.FailPipeline = clientStage.FailPipeline
	newStage.Job = clientStage.Job
	newStage.MarkUnstableAsSuccessful = clientStage.MarkUnstableAsSuccessful
	newStage.Master = clientStage.Master
	newStage.Parameters = clientStage.Parameters
	newStage.PropertyFile = clientStage.PropertyFile

	return stage(newStage)
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
