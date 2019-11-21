package provider

import "github.com/jgramoll/terraform-provider-spinnaker/client"

type trafficManagementOptions struct {
	EnableTraffic bool     `mapstructure:"enable_traffic"`
	Namespace     string   `mapstructure:"namespace"`
	Services      []string `mapstructure:"services"`
	Strategy      string   `mapstructure:"strategy"`
}

func newTrafficManagementOptions() *trafficManagementOptions {
	return &trafficManagementOptions{
		EnableTraffic: false,
	}
}

func toClientTrafficManagementOptions(options *[]*trafficManagementOptions) *client.TrafficManagementOptions {
	if options != nil {
		for _, o := range *options {
			newOptions := client.NewTrafficManagementOptions()
			newOptions.EnableTraffic = o.EnableTraffic
			newOptions.Namespace = o.Namespace
			newOptions.Services = o.Services
			newOptions.Strategy = o.Strategy
			return newOptions
		}
	}
	return nil
}

func fromClientTrafficManagementOptions(options *client.TrafficManagementOptions) *[]*trafficManagementOptions {
	var newArray []*trafficManagementOptions
	if options != nil {
		o := newTrafficManagementOptions()
		o.EnableTraffic = options.EnableTraffic
		o.Namespace = options.Namespace
		o.Services = options.Services
		o.Strategy = options.Strategy
		newArray = append(newArray, o)
	}
	return &newArray
}
