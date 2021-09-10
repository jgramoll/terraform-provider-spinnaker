package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var (
	testAccProviders map[string]terraform.ResourceProvider
	testAccProvider  *schema.Provider
)

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"spinnaker": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := testAccProvider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProviderConfigure(t *testing.T) {
	raw := map[string]interface{}{
		"address": "#address",
		"auth": map[string]interface{}{
			"cert_path": os.Getenv("SPINNAKER_CERT"),
			"key_path":  os.Getenv("SPINNAKER_KEY"),
		},
	}
	rawConfig, configErr := config.NewRawConfig(raw)
	if configErr != nil {
		t.Fatal(configErr)
	}
	c := terraform.NewResourceConfig(rawConfig)
	provider := Provider().(*schema.Provider)
	err := provider.Configure(c)
	if err != nil {
		t.Fatal(err)
	}

	testConfig := provider.Meta().(*Services).Config
	if testConfig.Address != raw["address"] {
		t.Fatalf("address should be %#v, not %#v", raw["address"], testConfig.Address)
	}

	auth, ok := raw["auth"].(map[string]interface{})
	if !ok {
		t.Fatal("auth is not present")
	}

	if testConfig.Auth.CertPath != auth["cert_path"] {
		t.Fatalf("certPath should be %#v, not %#v", auth["cert_path"], testConfig.Auth.CertPath)
	}
	if testConfig.Auth.KeyPath != auth["key_path"] {
		t.Fatalf("keyPath should be %#v, not %#v", auth["key_path"], testConfig.Auth.KeyPath)
	}
}

func testAccPreCheck(t *testing.T) {
	hasAuthCfg := os.Getenv("SPINNAKER_CERT") != "" && os.Getenv("SPINNAKER_KEY") != ""
	if !hasAuthCfg {
		t.Fatal("Spinnaker config (SPINNAKER_CERT and SPINNAKER_KEY) must be set for acceptance tests")
	}

	c := terraform.NewResourceConfig(nil)
	err := testAccProvider.Configure(c)
	if err != nil {
		t.Fatal(err)
	}
}
