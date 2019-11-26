package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func TestAccApplicationPipeline(t *testing.T) {
	var applicationRef client.Application
	name := fmt.Sprintf("tfacctest%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resourceName := "spinnaker_application.test"

	if os.Getenv("SKIP_APPLICATION_TEST") != "" {
		t.Skip()
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
					resource.TestCheckResourceAttr(resourceName, "permissions.0.execute.0", "spinnakeradmin"),
					resource.TestCheckResourceAttr(resourceName, "permissions.0.execute.1", "spinnakerexecute"),
					resource.TestCheckResourceAttr(resourceName, "permissions.0.read.0", "spinnakeradmin"),
					resource.TestCheckResourceAttr(resourceName, "permissions.0.read.1", "spinnakerread"),
					resource.TestCheckResourceAttr(resourceName, "permissions.0.write.0", "spinnakeradmin"),
					resource.TestCheckResourceAttr(resourceName, "permissions.0.write.1", "spinnakerwrite"),
				),
			},
		},
	})
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
	
	permissions {
		execute = [
			"spinnakeradmin",
			"spinnakerexecute"
		]
		read = [
			"spinnakeradmin",
			"spinnakerread"
		]
		write = [
			"spinnakeradmin",
			"spinnakerwrite"
		]
	}
}
`, name, name)
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
