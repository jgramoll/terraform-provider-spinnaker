package provider

import "github.com/get-bridge/terraform-provider-spinnaker/client"

type trafficManagement struct {
	Enabled bool                         `mapstructure:"enabled"`
	Options *[]*trafficManagementOptions `mapstructure:"options"`
}

func newTrafficManagement() *trafficManagement {
	return &trafficManagement{
		Enabled: false,
		Options: &[]*trafficManagementOptions{},
	}
}

func toClientTrafficManagement(trafficeManagement *[]*trafficManagement) *client.TrafficManagement {
	if trafficeManagement != nil {
		for _, t := range *trafficeManagement {
			newTM := client.NewTrafficManagement()
			newTM.Enabled = t.Enabled
			newTM.Options = toClientTrafficManagementOptions(t.Options)
			return newTM
		}
	}
	return nil
}

func fromClientTrafficManagement(clientTrafficManagement *client.TrafficManagement) *[]*trafficManagement {
	if clientTrafficManagement == nil {
		return nil
	}
	t := newTrafficManagement()
	t.Enabled = clientTrafficManagement.Enabled
	t.Options = fromClientTrafficManagementOptions(clientTrafficManagement.Options)
	newArray := []*trafficManagement{t}
	return &newArray
}
