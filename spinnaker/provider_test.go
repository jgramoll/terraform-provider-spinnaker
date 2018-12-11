package spinnaker

import (
  "testing"

  "github.com/hashicorp/terraform/config"
  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
)

var provider *schema.Provider
var raw map[string]interface{}
var rawConfig *config.RawConfig

func init() {
  provider = Provider()
  raw = map[string]interface{} {
    "address": "#address",
    "certPath": "#certPath",
    "keyPath": "#keyPath",
  }
  rawConfig, _ = config.NewRawConfig(raw)
}

func TestProviderConfigure(t *testing.T) {
  err := provider.Configure(terraform.NewResourceConfig(rawConfig))
  if err != nil {
    t.Fatalf("err: %s", err)
  }

  client := provider.Meta().(*Client)
  if client.Config.Address != raw["address"] {
    assertFail(t, "address", client.Config.Address)
  }
  if client.Config.CertPath != raw["certPath"] {
    assertFail(t, "certPath", client.Config.CertPath)
  }
  if client.Config.KeyPath != raw["keyPath"] {
    assertFail(t, "keyPath", client.Config.KeyPath)
  }
}

func assertFail(t *testing.T, field string, actual string) {
  t.Fatalf("address should be %#v, not %#v", raw[field], actual)
}
