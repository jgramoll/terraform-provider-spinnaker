package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type canaryAnalysisConfigScopes []*canaryAnalysisConfigScope

type canaryAnalysisConfigScope struct {
	ExtendedScopeParams map[string]string `mapstructure:"extended_scope_params"`
	ScopeName           string            `mapstructure:"scope_name"`
}

func (scopes *canaryAnalysisConfigScopes) toClientCanaryConfigScopes() *[]*client.CanaryAnalysisConfigScope {
	clientScopes := []*client.CanaryAnalysisConfigScope{}
	for _, c := range *scopes {
		clientScopes = append(clientScopes, &client.CanaryAnalysisConfigScope{
			ExtendedScopeParams: c.ExtendedScopeParams,
			ScopeName:           c.ScopeName,
		})
	}
	return &clientScopes
}

func (*canaryAnalysisConfigScopes) fromClientCanaryConfigScopes(clientScopes *[]*client.CanaryAnalysisConfigScope) *canaryAnalysisConfigScopes {
	scopes := canaryAnalysisConfigScopes{}
	for _, c := range *clientScopes {
		scopes = append(scopes, &canaryAnalysisConfigScope{
			ExtendedScopeParams: c.ExtendedScopeParams,
			ScopeName:           c.ScopeName,
		})
	}
	return &scopes
}
