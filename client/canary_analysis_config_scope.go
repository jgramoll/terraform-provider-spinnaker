package client

type CanaryAnalysisConfigScope struct {
	ControlLocation     string            `json:"controlLocation"`
	ControlScope        string            `json:"controlScope"`
	ExperimentLocation  string            `json:"experimentLocation"`
	ExperimentScope     string            `json:"experimentScope"`
	ExtendedScopeParams map[string]string `json:"extendedScopeParams"`
	ScopeName           string            `json:"scopeName"`
}
