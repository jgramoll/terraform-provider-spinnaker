package client

// StageExecutionWindowJitter random jitter to add to execution window
type StageExecutionWindowJitter struct {
	Enabled    bool `json:"enabled"`
	MaxDelay   int  `json:"maxDelay"`
	MinDelay   int  `json:"minDelay"`
	SkipManual bool `json:"skipManual"`
}

// StageExecutionWindowWhitelist which hours to deploy
type StageExecutionWindowWhitelist struct {
	EndHour   int `json:"endHour"`
	EndMin    int `json:"endMin"`
	StartHour int `json:"startHour"`
	StartMin  int `json:"startMin"`
}

// StageExecutionWindow when to execute pipeline stage
type StageExecutionWindow struct {
	Days      []int                             `json:"days"`
	Jitter    *StageExecutionWindowJitter       `json:"jitter"`
	Whitelist *[]*StageExecutionWindowWhitelist `json:"whitelist"`
}
