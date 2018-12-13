package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccKubernetesSecret_basic(t *testing.T) {
	var pipeline client.Pipeline
	app := "app"
	name := fmt.Sprintf("tf-acc-test-%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		// IDRefreshName: "spinnaker_pipeline.test",
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineConfigBasic(app, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists("spinnaker_pipeline.test", &pipeline),
					resource.TestCheckResourceAttr("spinnaker_pipeline.test", "name", name),
					resource.TestCheckResourceAttr("spinnaker_pipeline.test", "application", app),
					testAccCheckPipeline(&pipeline, map[string]string{"TestAnnotationOne": "one", "TestAnnotationTwo": "two"}),
				),
			},
			{
				Config: testAccPipelineConfigBasic(app, name+"-changed"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists("spinnaker_pipeline.test", &pipeline),
					resource.TestCheckResourceAttr("spinnaker_pipeline.test", "name", name+"-changed"),
					resource.TestCheckResourceAttr("spinnaker_pipeline.test", "application", app),
					testAccCheckPipeline(&pipeline, map[string]string{"TestAnnotationOne": "one", "TestAnnotationTwo": "two"}),
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
}
`, app, name)
}

func testAccCheckPipelineExists(resourceName string, outPipeline *client.Pipeline) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		c := testAccProvider.Meta().(*client.Client)
		pipeline, err := c.GetPipeline(rs.Primary.Attributes["application"], rs.Primary.Attributes["name"])
		if err != nil {
			return err
		}

		*outPipeline = *pipeline
		return nil
	}
}

func testAccCheckPipeline(pipeline *client.Pipeline, expected map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// 	if len(expected) == 0 && len(om.Annotations) == 0 {
		// 		return nil
		// 	}

		// 	// Remove any internal k8s annotations unless we expect them
		// 	annotations := om.Annotations
		// 	for key, _ := range annotations {
		// 		_, isExpected := expected[key]
		// 		if isInternalKey(key) && !isExpected {
		// 			delete(annotations, key)
		// 		}
		// 	}

		// 	if !reflect.DeepEqual(annotations, expected) {
		// 		return fmt.Errorf("%s annotations don't match.\nExpected: %q\nGiven: %q",
		// 			om.Name, expected, om.Annotations)
		// 	}
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
