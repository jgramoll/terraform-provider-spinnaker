package provider

import (
	"fmt"
	"os/user"
	"testing"

	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

var (
	testAccProviders map[string]terraform.ResourceProvider
	testAccProvider  *schema.Provider
	usr              *user.User
	raw              map[string]interface{}
)

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"spinnaker": testAccProvider,
	}
	usr, _ = user.Current()

	raw = map[string]interface{}{
		"address":   "#address",
		"cert_path": usr.HomeDir + "/.spin/client.crt",
		"key_path":  usr.HomeDir + "/.spin/client.key",
	}
}

func TestProvider(t *testing.T) {
	if err := testAccProvider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func TestProviderConfigure(t *testing.T) {
	testAccPreCheck(t)
	c := testAccProvider.Meta().(*client.Client)
	if c.Config.Address != raw["address"] {
		t.Fatalf("address should be %#v, not %#v", raw["address"], c.Config.Address)
	}
	if c.Config.CertPath != raw["cert_path"] {
		t.Fatalf("certPath should be %#v, not %#v", raw["cert_path"], c.Config.CertPath)
	}
	if c.Config.KeyPath != raw["key_path"] {
		t.Fatalf("keyPath should be %#v, not %#v", raw["key_path"], c.Config.KeyPath)
	}
}

func testAccPreCheck(t *testing.T) {
	rawConfig, configErr := config.NewRawConfig(raw)
	if configErr != nil {
		t.Fatal(configErr)
	}
	c := terraform.NewResourceConfig(rawConfig)
	err := testAccProvider.Configure(c)
	if err != nil {
		t.Fatal(err)
	}
}
