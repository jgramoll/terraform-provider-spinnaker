package provider

import (
        "log"

        "github.com/hashicorp/terraform/helper/schema"
        "github.com/jgramoll/terraform-provider-spinnaker/client"
        "github.com/mitchellh/mapstructure"
)

func Provider() *schema.Provider {
  return &schema.Provider{
    Schema: map[string]*schema.Schema{
      "address": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },

      "certPath": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },

      "keyPath": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
    },
    ResourcesMap: map[string]*schema.Resource{
      "spinnaker_pipeline": pipelineResource(),
    },
    ConfigureFunc: providerConfigure,
  }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
  var config client.Config
  configRaw := d.Get("").(map[string]interface{})
  if err := mapstructure.Decode(configRaw, &config); err != nil {
    return nil, err
  }

  log.Println("[INFO] Initializing Spinnaker client")
  return client.NewClient(config), nil
}
