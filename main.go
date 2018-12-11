package main

import (
  "github.com/hashicorp/terraform/plugin"
  "github.com/hashicorp/terraform/terraform"
  "github.com/jgramoll/terraform-provider-spinnaker/spinnaker"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: func() terraform.ResourceProvider {
      return spinnaker.Provider()
    },
  })
}
