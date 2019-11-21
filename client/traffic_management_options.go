package client

type TrafficManagementOptions struct {
	EnableTraffic bool     `json:"enableTraffic"`
	Namespace     string   `json:"namespace,omitempty"`
	Services      []string `json:"services,omitempty"`
	Strategy      string   `json:"strategy,omitempty"`
}

func NewTrafficManagementOptions() *TrafficManagementOptions {
	return &TrafficManagementOptions{
		EnableTraffic: false,
		Services:      []string{},
	}
}
