package provider

import "github.com/jgramoll/terraform-provider-spinnaker/client"

type trafficManagementOptions struct {
	EnableTraffic bool `mapstructure:"enable_traffic"`
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
		newArray = append(newArray, o)
	}
	return &newArray
}
