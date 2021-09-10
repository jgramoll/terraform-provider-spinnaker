package provider

import "github.com/get-bridge/terraform-provider-spinnaker/client"

type relationships struct {
	LoadBalancers  *[]string `mapstructure:"load_balancers"`
	SecurityGroups *[]string `mapstructure:"security_groups"`
}

func newRelationships() *relationships {
	return &relationships{
		LoadBalancers:  &[]string{},
		SecurityGroups: &[]string{},
	}
}

func toClientRelationships(relationships *[]*relationships) *client.Relationships {
	if relationships != nil {
		for _, r := range *relationships {
			if r != nil {
				newRelationships := client.Relationships(*r)
				return &newRelationships
			}
		}
	}
	return nil
}

func fromClientRelationships(clientRelationships *client.Relationships) *[]*relationships {
	if clientRelationships == nil {
		return nil
	}
	r := relationships(*clientRelationships)
	newRelationshipsArray := []*relationships{&r}
	return &newRelationshipsArray
}
