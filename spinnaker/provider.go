package spinnaker

import (
        "log"

        "github.com/hashicorp/terraform/helper/schema"
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
      "spinnaker_pipeline": pipeline(),
    },
    ConfigureFunc: providerConfigure,
  }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
  var config Config
  configRaw := d.Get("").(map[string]interface{})
  if err := mapstructure.Decode(configRaw, &config); err != nil {
    return nil, err
  }

  log.Println("[INFO] Initializing Spinnaker client")
  return NewClient(config), nil
}
