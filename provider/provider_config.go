package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

// config for provider
type providerConfig struct {
	Address string `mapstructure:"address"`
	// TODO I couldn't get EnvDefaultFunc working if nested with auth
	Enabled   bool   `mapstructure:"enabled"`
	CertPath  string `mapstructure:"cert_path"`
	KeyPath   string `mapstructure:"key_path"`
	UserEmail string `mapstructure:"user_email"`
	Insecure  bool   `mapstructure:"insecure"`
}

func newProviderConfig() *providerConfig {
	return &providerConfig{
		Enabled:  true,
		Insecure: true,
	}
}

func (c *providerConfig) toClientConfig() *client.Config {
	clientConfig := client.NewConfig()
	clientConfig.Address = c.Address
	clientConfig.Auth.Enabled = c.Enabled
	clientConfig.Auth.CertPath = c.CertPath
	clientConfig.Auth.KeyPath = c.KeyPath
	clientConfig.Auth.UserEmail = c.UserEmail
	clientConfig.Auth.Insecure = c.Insecure
	return clientConfig
}
