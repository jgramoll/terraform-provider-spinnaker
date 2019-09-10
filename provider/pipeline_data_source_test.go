package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccPipelineDataSourceBasic(t *testing.T) {
	app := "app"
	name := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineDataSourceConfigBasic(app, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.spinnaker_pipeline.test", "id"),
					resource.TestCheckResourceAttr("data.spinnaker_pipeline.test", "application", app),
					resource.TestCheckResourceAttr("data.spinnaker_pipeline.test", "name", name),
				),
			},
		},
	})
}

func testAccPipelineDataSourceConfigBasic(app string, name string) string {
	return testAccPipelineConfigBasic(app, name) + `
data "spinnaker_pipeline" "test" {
	application = "${spinnaker_pipeline.test.application}"
	name        = "${spinnaker_pipeline.test.name}"
}
`
}
