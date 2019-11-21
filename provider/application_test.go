package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func TestAccApplicationBasic(t *testing.T) {
	var applicationRef client.Application
	name := fmt.Sprintf("tfacctest%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	newName := name + "_changed"
	resourceName := "spinnaker_application.test"

	if os.Getenv("SKIP_APPLICATION_TEST") != "" {
		return
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckApplicationExists(resourceName, &applicationRef),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				ResourceName:  resourceName,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`403 Forbidden`),
			},
			{
				ResourceName:      resourceName,
				ImportStateId:     name,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccApplicationConfigBasic(newName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckApplicationExists(resourceName, &applicationRef),
					resource.TestCheckResourceAttr(resourceName, "name", newName),
				),
			},
		},
	})
}

func TestAccApplicationPipeline(t *testing.T) {
	var applicationRef client.Application
	name := acctest.RandStringFromCharSet(50, acctest.CharSetAlphaNum)
	newName := name + "_changed"
	resourceName := "spinnaker_application.test"
	pipeline := "spinnaker_pipeline.test"

	if os.Getenv("SKIP_APPLICATION_TEST") != "" {
		return
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationConfigPipeline(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckApplicationExists(resourceName, &applicationRef),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(pipeline, "application", name),
					resource.TestCheckResourceAttr(pipeline, "name", fmt.Sprintf("%s-pipeline", name)),
				),
			},
			{
				Config: testAccApplicationConfigPipeline(newName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckApplicationExists(resourceName, &applicationRef),
					resource.TestCheckResourceAttr(resourceName, "name", newName),
					resource.TestCheckResourceAttr(pipeline, "application", newName),
					resource.TestCheckResourceAttr(pipeline, "name", fmt.Sprintf("%s-pipeline", newName)),
				),
			},
		},
	})
}

func testAccApplicationConfigBasic(name string) string {
	return fmt.Sprintf(`
resource "spinnaker_application" "test" {
  name  		 = "%s"
  email 		 = "%s@spin.com"
  instance_port  = "8080"
  
  cloud_providers = [
	  "aws"
  ]
}`, name, name)
}

func testAccApplicationConfigPipeline(name string) string {
	return fmt.Sprintf(`
resource "spinnaker_application" "test" {
  name  		 = "%s"
  email 		 = "%s@spin.com"
  instance_port  = "8080"

  cloud_providers = [
	  "aws"
  ]
}

resource "spinnaker_pipeline" "test" {
	application = "%s"
	name        = "%s-pipeline"
	index       = 2
}

`, name, name, name, name)
}

func testAccCheckApplicationExists(resourceName string, a *client.Application) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		applicationService := testAccProvider.Meta().(*Services).ApplicationService
		app, err := applicationService.GetApplicationByName(rs.Primary.Attributes["name"])
		if err != nil {
			return err
		}
		*a = *app

		return nil
	}
}

func testAccCheckApplicationDestroy(s *terraform.State) error {
	applicationService := testAccProvider.Meta().(*Services).ApplicationService
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "spinnaker_application" {
			_, err := applicationService.GetApplicationByName(rs.Primary.Attributes["name"])
			if err == nil {
				return fmt.Errorf("Application still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
