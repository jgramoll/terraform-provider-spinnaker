package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccPipeline_basic(t *testing.T) {
	app := "app"
	name := fmt.Sprintf("tf-acc-test-%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineConfigBasic(app, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists("spinnaker_pipeline.test"),
					resource.TestCheckResourceAttr("spinnaker_pipeline.test", "name", name),
					resource.TestCheckResourceAttr("spinnaker_pipeline.test", "application", app),
				),
			},
			{
				Config: testAccPipelineConfigBasic(app, name+"-changed"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists("spinnaker_pipeline.test"),
					resource.TestCheckResourceAttr("spinnaker_pipeline.test", "name", name+"-changed"),
					resource.TestCheckResourceAttr("spinnaker_pipeline.test", "application", app),
				),
			},
		},
	})
}

func testAccPipelineConfigBasic(app string, name string) string {
	return fmt.Sprintf(`
resource "spinnaker_pipeline" "test" {
	application = "%s"
	name        = "%s"
	index       = 2
}
`, app, name)
}

func testAccCheckPipelineExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		c := testAccProvider.Meta().(*client.Client)
		_, err := c.GetPipeline(rs.Primary.Attributes["application"], rs.Primary.Attributes["name"])
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckPipelineDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "spinnaker_pipeline" {
			_, err := c.GetPipeline(rs.Primary.Attributes["application"], rs.Primary.Attributes["name"])
			if err == nil {
				return fmt.Errorf("Pipeline still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
