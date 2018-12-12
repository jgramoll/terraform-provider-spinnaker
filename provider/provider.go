package provider

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

// Config for provider
type Config struct {
	Address  string
	CertPath string `mapstructure:"cert_path"`
	KeyPath  string `mapstructure:"key_path"`
}

// Provider for terraform
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type: schema.TypeString,
				// Required: true,
				Optional: true,
			},

			"cert_path": &schema.Schema{
				Type: schema.TypeString,
				// Required: true,
				Optional: true,
			},

			"key_path": &schema.Schema{
				Type: schema.TypeString,
				// Required: true,
				Optional: true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"spinnaker_pipeline": pipelineResource(),
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
	return client.NewClient(client.Config(config)), nil
}
