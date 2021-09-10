package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

func fromClientCapacity(clientCapacity *client.Capacity) *[]*capacity {
	if clientCapacity == nil {
		return nil
	}
	newCapacity := capacity(*clientCapacity)
	return &[]*capacity{&newCapacity}
}

func toClientCapacity(c *[]*capacity) *client.Capacity {
	if c != nil {
		for _, c := range *c {
			newCapacity := client.Capacity(*c)
			return &newCapacity
		}
	}
	return nil
}
