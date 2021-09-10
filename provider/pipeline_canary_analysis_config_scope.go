package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type canaryAnalysisConfigScopes []*canaryAnalysisConfigScope

type canaryAnalysisConfigScope struct {
	ControlLocation     string            `mapstructure:"control_location"`
	ControlScope        string            `mapstructure:"control_scope"`
	ExperimentLocation  string            `mapstructure:"experiment_location"`
	ExperimentScope     string            `mapstructure:"experiment_scope"`
	ExtendedScopeParams map[string]string `mapstructure:"extended_scope_params"`
	ScopeName           string            `mapstructure:"scope_name"`
	Step                int               `mapstructure:"step"`
}

func (scopes *canaryAnalysisConfigScopes) toClientCanaryConfigScopes() *[]*client.CanaryAnalysisConfigScope {
	clientScopes := []*client.CanaryAnalysisConfigScope{}
	for _, c := range *scopes {
		clientScopes = append(clientScopes, &client.CanaryAnalysisConfigScope{
			ControlLocation:     c.ControlLocation,
			ControlScope:        c.ControlScope,
			ExperimentLocation:  c.ExperimentLocation,
			ExperimentScope:     c.ExperimentScope,
			ExtendedScopeParams: c.ExtendedScopeParams,
			ScopeName:           c.ScopeName,
			Step:                c.Step,
		})
	}
	return &clientScopes
}

func (*canaryAnalysisConfigScopes) fromClientCanaryConfigScopes(clientScopes *[]*client.CanaryAnalysisConfigScope) *canaryAnalysisConfigScopes {
	scopes := canaryAnalysisConfigScopes{}
	for _, c := range *clientScopes {
		scopes = append(scopes, &canaryAnalysisConfigScope{
			ControlLocation:     c.ControlLocation,
			ControlScope:        c.ControlScope,
			ExperimentLocation:  c.ExperimentLocation,
			ExperimentScope:     c.ExperimentScope,
			ExtendedScopeParams: c.ExtendedScopeParams,
			ScopeName:           c.ScopeName,
			Step:                c.Step,
		})
	}
	return &scopes
}
