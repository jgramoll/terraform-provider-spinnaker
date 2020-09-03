package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type deployManifestStage struct {
	baseStage `mapstructure:",squash"`

	Account                  string                `mapstructure:"account"`
	Credentials              string                `mapstructure:"credentials"`
	NamespaceOverride        string                `mapstructure:"namespace_override"`
	CloudProvider            string                `mapstructure:"cloud_provider"`
	ManifestArtifactAccount  string                `mapstructure:"manifest_artifact_account"`
	ManifestArtifactID       string                `mapstructure:"manifest_artifact_id"`
	Manifests                *manifests            `mapstructure:"manifests"`
	Moniker                  *[]*moniker           `mapstructure:"moniker"`
	Relationships            *[]*relationships     `mapstructure:"relationships"`
	SkipExpressionEvaluation bool                  `mapstructure:"skip_expression_evaluation"`
	Source                   string                `mapstructure:"source"`
	TrafficManagement        *[]*trafficManagement `mapstructure:"traffic_management"`
}

func newDeployManifestStage() *deployManifestStage {
	return &deployManifestStage{
		baseStage:         *newBaseStage(),
		Relationships:     &[]*relationships{},
		TrafficManagement: &[]*trafficManagement{},
	}
}

func (s *deployManifestStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewDeployManifestStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.Account = s.Account
	cs.Credentials = s.Credentials
	cs.NamespaceOverride = s.NamespaceOverride
	cs.CloudProvider = s.CloudProvider
	cs.ManifestArtifactAccount = s.ManifestArtifactAccount
	cs.ManifestArtifactID = s.ManifestArtifactID
	cs.Manifests = s.Manifests.toClientManifests()
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

func (*deployManifestStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.DeployManifestStage)
	newStage := newDeployManifestStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.Account = clientStage.Account
	newStage.Credentials = clientStage.Credentials
	newStage.NamespaceOverride = clientStage.NamespaceOverride
	newStage.CloudProvider = clientStage.CloudProvider
	newStage.ManifestArtifactAccount = clientStage.ManifestArtifactAccount
	newStage.ManifestArtifactID = clientStage.ManifestArtifactID
	newStage.Manifests = fromClientManifests(clientStage.Manifests)
	newStage.Moniker = fromClientMoniker(clientStage.Moniker)
	newStage.Relationships = fromClientRelationships(clientStage.Relationships)
	newStage.SkipExpressionEvaluation = clientStage.SkipExpressionEvaluation
	newStage.Source = clientStage.Source.String()
	newStage.TrafficManagement = fromClientTrafficManagement(clientStage.TrafficManagement)

	return newStage, nil
}

func (s *deployManifestStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("account", s.Account)
	if err != nil {
		return err
	}
	err = d.Set("credentials", s.Credentials)
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
	err = d.Set("manifest_artifact_id", s.ManifestArtifactID)
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
