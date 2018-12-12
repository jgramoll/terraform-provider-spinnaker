package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jgramoll/terraform-provider-spinnaker/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider
	})
}
