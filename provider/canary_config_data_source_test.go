package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCanaryConfigDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCanaryConfigDataSourceConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.spinnaker_canary_config.test", "id"),
					resource.TestCheckResourceAttr("data.spinnaker_canary_config.test", "name", name),
				),
			},
		},
	})
}

func testAccCanaryConfigDataSourceConfigBasic(name string) string {
	return testAccCanaryConfigConfigBasic(name) + `
data "spinnaker_canary_config" "test" {
	name        = "${spinnaker_canary_config.test.name}"
}
`
}
