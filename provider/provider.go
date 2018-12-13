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
	Address   string
	CertPath  string `mapstructure:"cert_path"`
	KeyPath   string `mapstructure:"key_path"`
	UserEmail string `mapstructure:"user_email"`
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

			"user_email": &schema.Schema{
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

	// debug.PrintStack()
	// fmt.Println("config", config)
	// fmt.Println("config", client.Config(config))
	log.Println("[INFO] Initializing Spinnaker client")

	// Why do we need this....
	config.Address = "https://api.spinnaker.inseng.net"
	config.CertPath = "/Users/jgramoll/.spin/client.crt"
	config.KeyPath = "/Users/jgramoll/.spin/client.key"
	config.UserEmail = "jgramoll@instructure.com"

	return client.NewClient(client.Config(config)), nil
}
