package client

// CanaryAnalysisConfigScope canary analysis
type CanaryAnalysisConfigScope struct {
	ControlLocation     string            `json:"controlLocation,omitempty"`
	ControlScope        string            `json:"controlScope,omitempty"`
	ExperimentLocation  string            `json:"experimentLocation,omitempty"`
	ExperimentScope     string            `json:"experimentScope,omitempty"`
	ExtendedScopeParams map[string]string `json:"extendedScopeParams"`
	ScopeName           string            `json:"scopeName"`
	Step                int               `json:"step"`
}
