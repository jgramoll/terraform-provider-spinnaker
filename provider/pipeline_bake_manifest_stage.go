package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type manifestInputArtifact struct {
	Account string `json:"account"`
	ID      string `json:"id"`
}

type bakeManifestStage struct {
	baseStage `mapstructure:",squash"`

	EvaluateOverrideExpressions bool `mapstructure:"evaluate_override_expressions"`

	InputArtifacts   []manifestInputArtifact `mapstructure:"input_artifact"`
	Namespace        string                  `mapstructure:"namespace"`
	OutputName       string                  `mapstructure:"output_name"`
	Overrides        map[string]string       `mapstructure:"overrides"`
	RawOverrides     bool                    `mapstructure:"raw_overrides"`
	TemplateRenderer string                  `mapstructure:"template_renderer"`
}

func newBakeManifestStage() *bakeManifestStage {
	return &bakeManifestStage{
		baseStage: *newBaseStage(),
	}
}

func (s *bakeManifestStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewBakeManifestStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.EvaluateOverrideExpressions = s.EvaluateOverrideExpressions
	for _, a := range s.InputArtifacts {
		cs.InputArtifacts = append(cs.InputArtifacts, client.ManifestInputArtifact(a))
	}
	cs.Namespace = s.Namespace
	cs.OutputName = s.OutputName
	cs.Overrides = s.Overrides
	cs.RawOverrides = s.RawOverrides
	cs.TemplateRenderer = s.TemplateRenderer

	return cs, nil
}

func (*bakeManifestStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.BakeManifestStage)
	newStage := newBakeManifestStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.EvaluateOverrideExpressions = clientStage.EvaluateOverrideExpressions
	for _, a := range clientStage.InputArtifacts {
		newStage.InputArtifacts = append(newStage.InputArtifacts, manifestInputArtifact(a))
	}
	newStage.Namespace = clientStage.Namespace
	newStage.OutputName = clientStage.OutputName
	newStage.Overrides = clientStage.Overrides
	newStage.RawOverrides = clientStage.RawOverrides
	newStage.TemplateRenderer = clientStage.TemplateRenderer

	return newStage, nil
}

func (s *bakeManifestStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("evaluate_override_expressions", s.EvaluateOverrideExpressions)
	if err != nil {
		return err
	}
	err = d.Set("input_artifact", s.InputArtifacts)
	if err != nil {
		return err
	}
	err = d.Set("namespace", s.Namespace)
	if err != nil {
		return err
	}
	err = d.Set("output_name", s.OutputName)
	if err != nil {
		return err
	}
	err = d.Set("overrides", s.Overrides)
	if err != nil {
		return err
	}
	err = d.Set("raw_overrides", s.RawOverrides)
	if err != nil {
		return err
	}
	err = d.Set("template_renderer", s.TemplateRenderer)
	if err != nil {
		return err
	}
	return nil
}
