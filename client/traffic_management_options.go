package client

type TrafficManagementOptions struct {
	EnableTraffic bool     `json:"enableTraffic"`
	Namespace     string   `json:"namespace"`
	Services      []string `json:"services"`
	Strategy      string   `json:"strategy"`
}

func NewTrafficManagementOptions() *TrafficManagementOptions {
	return &TrafficManagementOptions{
		EnableTraffic: false,
		Services:      []string{},
	}
}
