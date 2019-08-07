package client

type TrafficManagementOptions struct {
	EnableTraffic bool `json:"enableTraffic"`
}

func NewTrafficManagementOptions() *TrafficManagementOptions {
	return &TrafficManagementOptions{
		EnableTraffic: false,
	}
}
