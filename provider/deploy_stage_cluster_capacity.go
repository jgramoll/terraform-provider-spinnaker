package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type deployStageClusterCapacity struct {
	Desired int `mapstructure:"desired"`
	Max     int `mapstructure:"max"`
	Min     int `mapstructure:"min"`
}

// newCluster.Capacity = fromClientClusterCapacity(c.Capacity)

func fromClientClusterCapacity(clientCapacity *client.DeployStageClusterCapacity) *[]*deployStageClusterCapacity {
	if clientCapacity == nil {
		return nil
	}
	newCapacity := deployStageClusterCapacity(*clientCapacity)
	return &[]*deployStageClusterCapacity{&newCapacity}
}

func toClientClusterCapacity(capacity *[]*deployStageClusterCapacity) *client.DeployStageClusterCapacity {
	if capacity != nil {
		for _, c := range *capacity {
			newCapacity := client.DeployStageClusterCapacity(*c)
			return &newCapacity
		}
	}
	return nil
}
