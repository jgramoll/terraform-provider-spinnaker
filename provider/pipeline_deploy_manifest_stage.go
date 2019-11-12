package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type deployManifestStage struct {
	// baseStage
	Name                              string                   `mapstructure:"name"`
	RefID                             string                   `mapstructure:"ref_id"`
	Type                              client.StageType         `mapstructure:"type"`
	RequisiteStageRefIds              []string                 `mapstructure:"requisite_stage_ref_ids"`
	Notifications                     *[]*notification         `mapstructure:"notification"`
	StageEnabled                      *[]*stageEnabled         `mapstructure:"stage_enabled"`
	CompleteOtherBranchesThenFail     bool                     `mapstructure:"complete_other_branches_then_fail"`
	ContinuePipeline                  bool                     `mapstructure:"continue_pipeline"`
	FailOnFailedExpressions           bool                     `mapstructure:"fail_on_failed_expressions"`
	FailPipeline                      bool                     `mapstructure:"fail_pipeline"`
	OverrideTimeout                   bool                     `mapstructure:"override_timeout"`
	RestrictExecutionDuringTimeWindow bool                     `mapstructure:"restrict_execution_during_time_window"`
	RestrictedExecutionWindow         *[]*stageExecutionWindow `mapstructure:"restricted_execution_window"`
	// End baseStage

	Account                  string                `mapstructure:"account"`
	NamespaceOverride        string                `mapstructure:"namespace_override"`
	CloudProvider            string                `mapstructure:"cloud_provider"`
	ManifestArtifactAccount  string                `mapstructure:"manifest_artifact_account"`
	Manifests                *deployManifests      `mapstructure:"manifests"`
	Moniker                  *[]*moniker           `mapstructure:"moniker"`
	Relationships            *[]*relationships     `mapstructure:"relationships"`
	SkipExpressionEvaluation bool                  `mapstructure:"skip_expression_evaluation"`
	Source                   string                `mapstructure:"source"`
	TrafficManagement        *[]*trafficManagement `mapstructure:"traffic_management"`
}

func newDeployManifestStage() *deployManifestStage {
	return &deployManifestStage{
		Type:                    client.DeployManifestStageType,
		FailPipeline:            true,
		ManifestArtifactAccount: "docker-registry",
		Relationships:           &[]*relationships{},
		TrafficManagement:       &[]*trafficManagement{},
	}
}

func (s *deployManifestStage) toClientStage(config *client.Config) (client.Stage, error) {
	// baseStage
	notifications, err := toClientNotifications(s.Notifications)
	if err != nil {
		return nil, err
	}

	cs := client.NewDeployManifestStage()
	cs.Name = s.Name
	cs.RefID = s.RefID
	cs.RequisiteStageRefIds = s.RequisiteStageRefIds
	cs.Notifications = notifications
	cs.SendNotifications = notifications != nil && len(*notifications) > 0
	cs.StageEnabled = toClientStageEnabled(s.StageEnabled)
	cs.CompleteOtherBranchesThenFail = s.CompleteOtherBranchesThenFail
	cs.ContinuePipeline = s.ContinuePipeline
	cs.FailOnFailedExpressions = s.FailOnFailedExpressions
	cs.FailPipeline = s.FailPipeline
	cs.OverrideTimeout = s.OverrideTimeout
	cs.RestrictExecutionDuringTimeWindow = s.RestrictExecutionDuringTimeWindow
	cs.RestrictedExecutionWindow = toClientExecutionWindow(s.RestrictedExecutionWindow)
	// End baseStage

	cs.Account = s.Account
	cs.NamespaceOverride = s.NamespaceOverride
	cs.CloudProvider = s.CloudProvider
	cs.ManifestArtifactAccount = s.ManifestArtifactAccount
	cs.Manifests = s.Manifests.toClientDeployManifests()
	cs.Moniker = toClientMoniker(s.Moniker)
	cs.Relationships = toClientRelationships(s.Relationships)
	cs.SkipExpressionEvaluation = s.SkipExpressionEvaluation
	source, err := client.ParseDeployManifestSource(s.Source)
	if err != nil {
		return nil, err
	}
	cs.Source = source
	cs.TrafficManagement = toClientTrafficManagement(s.TrafficManagement)

	return cs, nil
}

func (s *deployManifestStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.DeployManifestStage)
	newStage := newDeployManifestStage()

	// baseStage
	newStage.Name = clientStage.Name
	newStage.RefID = clientStage.RefID
	newStage.RequisiteStageRefIds = clientStage.RequisiteStageRefIds
	newStage.Notifications = fromClientNotifications(clientStage.Notifications)
	newStage.StageEnabled = fromClientStageEnabled(clientStage.StageEnabled)
	newStage.CompleteOtherBranchesThenFail = clientStage.CompleteOtherBranchesThenFail
	newStage.ContinuePipeline = clientStage.ContinuePipeline
	newStage.FailOnFailedExpressions = clientStage.FailOnFailedExpressions
	newStage.FailPipeline = clientStage.FailPipeline
	newStage.OverrideTimeout = clientStage.OverrideTimeout
	newStage.RestrictExecutionDuringTimeWindow = clientStage.RestrictExecutionDuringTimeWindow
	newStage.RestrictedExecutionWindow = fromClientExecutionWindow(clientStage.RestrictedExecutionWindow)
	// end baseStage

	newStage.Account = clientStage.Account
	newStage.NamespaceOverride = clientStage.NamespaceOverride
	newStage.CloudProvider = clientStage.CloudProvider
	newStage.ManifestArtifactAccount = clientStage.ManifestArtifactAccount
	newStage.Manifests = fromClientDeployManifests(clientStage.Manifests)
	newStage.Moniker = fromClientMoniker(clientStage.Moniker)
	newStage.Relationships = fromClientRelationships(clientStage.Relationships)
	newStage.SkipExpressionEvaluation = clientStage.SkipExpressionEvaluation
	newStage.Source = clientStage.Source.String()
	newStage.TrafficManagement = fromClientTrafficManagement(clientStage.TrafficManagement)

	return newStage
}

func (s *deployManifestStage) SetResourceData(d *schema.ResourceData) error {
	// baseStage
	err := d.Set("name", s.Name)
	if err != nil {
		return err
	}
	err = d.Set("requisite_stage_ref_ids", s.RequisiteStageRefIds)
	if err != nil {
		return err
	}
	err = d.Set("notification", s.Notifications)
	if err != nil {
		return err
	}
	err = d.Set("stage_enabled", s.StageEnabled)
	if err != nil {
		return err
	}
	err = d.Set("complete_other_branches_then_fail", s.CompleteOtherBranchesThenFail)
	if err != nil {
		return err
	}
	err = d.Set("continue_pipeline", s.ContinuePipeline)
	if err != nil {
		return err
	}
	err = d.Set("fail_on_failed_expressions", s.FailOnFailedExpressions)
	if err != nil {
		return err
	}
	err = d.Set("fail_pipeline", s.FailPipeline)
	if err != nil {
		return err
	}
	err = d.Set("override_timeout", s.OverrideTimeout)
	if err != nil {
		return err
	}
	err = d.Set("restrict_execution_during_time_window", s.RestrictExecutionDuringTimeWindow)
	if err != nil {
		return err
	}
	err = d.Set("restricted_execution_window", s.RestrictedExecutionWindow)
	if err != nil {
		return err
	}
	// End baseStage

	err = d.Set("account", s.Account)
	if err != nil {
		return err
	}
	err = d.Set("namespace_override", s.NamespaceOverride)
	if err != nil {
		return err
	}
	err = d.Set("cloud_provider", s.CloudProvider)
	if err != nil {
		return err
	}
	err = d.Set("manifest_artifact_account", s.ManifestArtifactAccount)
	if err != nil {
		return err
	}
	err = d.Set("manifests", s.Manifests)
	if err != nil {
		return err
	}
	err = d.Set("moniker", s.Moniker)
	if err != nil {
		return err
	}
	err = d.Set("relationships", s.Relationships)
	if err != nil {
		return err
	}
	err = d.Set("skip_expression_evaluation", s.SkipExpressionEvaluation)
	if err != nil {
		return err
	}
	err = d.Set("source", s.Source)
	if err != nil {
		return err
	}
	return d.Set("traffic_management", s.TrafficManagement)
}

func (s *deployManifestStage) SetRefID(id string) {
	s.RefID = id
}

func (s *deployManifestStage) GetRefID() string {
	return s.RefID
}
