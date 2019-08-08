package provider

import "github.com/jgramoll/terraform-provider-spinnaker/client"

type trafficManagement struct {
	Enabled bool                      `mapstructure:"enabled"`
	Options *trafficManagementOptions `mapstructure:"options"`
}

func newTrafficManagement() *trafficManagement {
	return &trafficManagement{
		Enabled: false,
		Options: newTrafficManagementOptions(),
	}
}

func toClientTrafficManagement(m *[]*trafficManagement) *client.TrafficManagement {
	if m != nil {
		for _, t := range *m {
			newTM := client.NewTrafficManagement()
			newTM.Enabled = t.Enabled
			*newTM.Options = client.TrafficManagementOptions(*t.Options)
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
	*t.Options = trafficManagementOptions(*clientTrafficManagement.Options)
	newArray := []*trafficManagement{t}
	return &newArray
}
