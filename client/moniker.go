package client

// Moniker for cluster
type Moniker struct {
	App     string `json:"app,omitempty"`
	Cluster string `json:"cluster,omitempty"`
	Detail  string `json:"detail,omitempty"`
	Stack   string `json:"stack,omitempty"`
}
