package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func TestAccCanaryConfigBasic(t *testing.T) {
	var canaryConfigRef client.CanaryConfig
	name := fmt.Sprintf("tfacctest%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	newName := name + "_changed"
	resourceName := "spinnaker_canary_config.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCanaryConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCanaryConfigConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCanaryConfigExists(resourceName, &canaryConfigRef),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				ResourceName:  resourceName,
				ImportStateId: "non-existent",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Cannot import non-existent remote object`),
			},
			{
				ResourceName:      resourceName,
				ImportStateId:     canaryConfigRef.ID,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCanaryConfigConfigBasic(newName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCanaryConfigExists(resourceName, &canaryConfigRef),
					resource.TestCheckResourceAttr(resourceName, "name", newName),
				),
			},
		},
	})
}

func testAccCanaryConfigConfigBasic(name string) string {
	return fmt.Sprintf(`
resource "spinnaker_canary_config" "test" {
	name         = "%s"
	applications = ["app"]
	metric {
		groups = ["Group 1"]
		name = "my metric"
		query {
			metric_name = "avg:aws.ec2.cpucredit_balance"
			service_type = "datadog"
			type = "datadog"
		}
	}
	classifier {
		group_weights = {
			"Group 1" = 100
		}
	}
	judge {
		name = "NetflixACAJudge-v1.0"
	}
}`, name)
}

func testAccCheckCanaryConfigExists(resourceName string, a *client.CanaryConfig) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		canaryConfigService := testAccProvider.Meta().(*Services).CanaryConfigService
		app, err := canaryConfigService.GetCanaryConfig(rs.Primary.ID)
		if err != nil {
			return err
		}
		*a = *app

		return nil
	}
}

func testAccCheckCanaryConfigDestroy(s *terraform.State) error {
	canaryConfigService := testAccProvider.Meta().(*Services).CanaryConfigService
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "spinnaker_canary_config" {
			_, err := canaryConfigService.GetCanaryConfig(rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Canary Config still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
