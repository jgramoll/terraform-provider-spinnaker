package provider

import "github.com/get-bridge/terraform-provider-spinnaker/client"

type applicationPermissions struct {
	Execute []string `mapstructure:"execute"`
	Read    []string `mapstructure:"read"`
	Write   []string `mapstructure:"write"`
}

func newApplicationPermission() *applicationPermissions {
	return &applicationPermissions{}
}

func toClientApplicationPermissions(p *[]applicationPermissions) *client.ApplicationPermissions {
	if len(*p) == 0 {
		return nil
	}
	permissions := (*p)[0]

	clientPermission := client.NewApplicationPermissions()
	clientPermission.Execute = permissions.Execute
	clientPermission.Read = permissions.Read
	clientPermission.Write = permissions.Write
	return clientPermission
}

func fromClientApplicationPermissions(clientPermission *client.ApplicationPermissions) *[]applicationPermissions {
	if clientPermission == nil {
		return nil
	}
	p := newApplicationPermission()

	p.Execute = clientPermission.Execute
	p.Read = clientPermission.Read
	p.Write = clientPermission.Write

	return &[]applicationPermissions{*p}
}
