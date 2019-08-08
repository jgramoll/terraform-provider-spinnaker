package provider

type trafficManagementOptions struct {
	EnableTraffic bool `mapstructure:"enable_traffic"`
}

func newTrafficManagementOptions() *trafficManagementOptions {
	return &trafficManagementOptions{
		EnableTraffic: false,
	}
}
