package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func TestAccApplicationBasic(t *testing.T) {
	var applicationRef client.Application
	name := fmt.Sprintf("tf-acc-test-%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	newName := name + "-changed"
	resourceName := "spinnaker_application.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckApplicationExists(resourceName, &applicationRef),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
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

func TestAccApplicationTrigger(t *testing.T) {
	var applicationRef client.Application
	app := "app"
	name := fmt.Sprintf("tf-acc-test-%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	newName := name + "-changed"
	resourceName := "spinnaker_application.test"
	pipeline := "spinnaker_pipeline.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineConfigTrigger(app, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckApplicationExists(resourceName, &applicationRef),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(pipeline, "application", name),
					resource.TestCheckResourceAttr(pipeline, "name", fmt.Sprintf("%s-pipe", name)),
				),
			},
			{
				Config: testAccPipelineConfigTrigger(app, newName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckApplicationExists(resourceName, &applicationRef),
					resource.TestCheckResourceAttr(resourceName, "name", newName),
					resource.TestCheckResourceAttr(pipeline, "application", newName),
					resource.TestCheckResourceAttr(pipeline, "name", fmt.Sprintf("%s-pipe", newName)),
				),
			},
		},
	})
}

func testAccApplicationConfigBasic(app string) string {
	return fmt.Sprintf(`
resource "spinnaker_application" "test" {
  name  = "%s"
  email = "%s@%s.com
}`, app, app, app)
}

func testAccApplicationConfigTrigger(app string, name string) string {
	return fmt.Sprintf(`
resource "spinnaker_application" "test" {
  name  = "%s"
  email = "%s@%s.com
}

resource "spinnaker_pipeline" "test" {
	application = "%s"
	name        = "%s-pipeline"
	index       = 2
}

`, app, app, app, app, name)
}

func testAccCheckApplicationExists(resourceName string, a *client.Application) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		applicationService := testAccProvider.Meta().(*Services).ApplicationService
		app, err := applicationService.GetApplicationByName(rs.Primary.Attributes["application"])
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
			_, err := applicationService.GetApplicationByName(rs.Primary.Attributes["application"])
			if err == nil {
				return fmt.Errorf("Application still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
