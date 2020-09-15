package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type moniker struct {
	App      string `mapstructure:"app"`
	Cluster  string `mapstructure:"cluster"`
	Detail   string `mapstructure:"detail"`
	Stack    string `mapstructure:"stack"`
	Sequence string `mapstructure:"sequence"`
}

func fromClientMoniker(clientMoniker *client.Moniker) *[]*moniker {
	if clientMoniker == nil {
		return nil
	}
	newMoniker := moniker(*clientMoniker)
	newMonikerArray := []*moniker{&newMoniker}
	return &newMonikerArray
}

func toClientMoniker(moniker *[]*moniker) *client.Moniker {
	if moniker != nil && len(*moniker) > 0 {
		for _, m := range *moniker {
			if m == nil {
				return nil
			}
			newMoniker := client.Moniker(*m)
			return &newMoniker
		}
	}
	return nil
}

// func (c *deployStageCluster) clientMoniker() client.Moniker {
// 	if len(c.Moniker) > 0 {
// 		return client.Moniker(c.Moniker[0])
// 	}
// 	return client.Moniker{}
// }
