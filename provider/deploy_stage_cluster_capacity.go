package provider

import (
	"strconv"

	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func fromClientCapacity(clientCapacity *client.Capacity) *[]*capacity {
	if clientCapacity == nil {
		return nil
	}
	newCapacity := capacity{}
	if n, err := strconv.Atoi(clientCapacity.Max); err == nil {
		newCapacity.Max = n
	} else {
		newCapacity.MaxExpression = clientCapacity.Max
	}
	if n, err := strconv.Atoi(clientCapacity.Min); err == nil {
		newCapacity.Min = n
	} else {
		newCapacity.MinExpression = clientCapacity.Min
	}
	if n, err := strconv.Atoi(clientCapacity.Desired); err == nil {
		newCapacity.Desired = n
	} else {
		newCapacity.DesiredExpression = clientCapacity.Desired
	}
	return &[]*capacity{&newCapacity}
}

func toClientCapacity(c *[]*capacity) *client.Capacity {
	if c != nil {
		for _, c := range *c {
			newCapacity := client.Capacity{}
			if c.MaxExpression != "" {
				newCapacity.Max = c.MaxExpression
			} else {
				newCapacity.Max = strconv.Itoa(c.Max)
			}
			if c.MinExpression != "" {
				newCapacity.Min = c.MinExpression
			} else {
				newCapacity.Min = strconv.Itoa(c.Min)
			}
			if c.DesiredExpression != "" {
				newCapacity.Desired = c.DesiredExpression
			} else {
				newCapacity.Desired = strconv.Itoa(c.Desired)
			}
			return &newCapacity
		}
	}
	return nil
}
