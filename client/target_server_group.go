package client

// TargetServerGroupStage for pipeline
type TargetServerGroupStage struct {
	CloudProvider     string   `json:"cloudProvider"`
	CloudProviderType string   `json:"cloudProviderType"`
	Cluster           string   `json:"cluster"`
	Credentials       string   `json:"credentials"`
	Moniker           *Moniker `json:"moniker"`
	Regions           []string `json:"regions"`
	Target            string   `json:"target"`
}

func newTargetServerGroupStage() *TargetServerGroupStage {
	return &TargetServerGroupStage{}
}
