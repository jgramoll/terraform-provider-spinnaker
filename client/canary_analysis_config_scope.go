package client

type CanaryAnalysisConfigScope struct {
	ExtendedScopeParams map[string]string `json:"extendedScopeParams"`
	ScopeName           string            `json:"scopeName"`
}
