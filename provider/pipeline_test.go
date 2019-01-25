package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func TestAccPipelineBasic(t *testing.T) {
	app := "app"
	name := fmt.Sprintf("tf-acc-test-%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	newName := name + "-changed"
	pipeline := "spinnaker_pipeline.test"
	pipelineParameterName := fmt.Sprintf("My %s parameter", name)
	pipelineParameterNewName := fmt.Sprintf("My %s parameter", newName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineConfigBasic(app, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipeline),
					resource.TestCheckResourceAttr(pipeline, "name", name),
					resource.TestCheckResourceAttr(pipeline, "parameter.0.name", pipelineParameterName),
					resource.TestCheckResourceAttr(pipeline, "parameter.1.default", "mosdef"),
					resource.TestCheckResourceAttr(pipeline, "parameter.1.label", "whatevs"),
					resource.TestCheckResourceAttr(pipeline, "application", app),
					testAccCheckPipelineParameters(pipeline, []string{pipelineParameterName, "Detailed parameter"}),
				),
			},
			{
				Config: testAccPipelineConfigBasic(app, newName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipeline),
					resource.TestCheckResourceAttr(pipeline, "name", newName),
					resource.TestCheckResourceAttr(pipeline, "parameter.0.name", pipelineParameterNewName),
					resource.TestCheckResourceAttr(pipeline, "parameter.1.default", "mosdef"),
					resource.TestCheckResourceAttr(pipeline, "parameter.1.label", "whatevs"),
					resource.TestCheckResourceAttr(pipeline, "application", app),
					testAccCheckPipelineParameters(pipeline, []string{pipelineParameterNewName, "Detailed parameter"}),
				),
			},
		},
	})
}

func TestAccPipelineTrigger(t *testing.T) {
	app := "app"
	name := fmt.Sprintf("tf-acc-test-%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	newName := name + "-changed"
	pipeline := "spinnaker_pipeline.test"
	trigger := "spinnaker_pipeline_trigger.jenkins"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineConfigTrigger(app, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipeline),
					resource.TestCheckResourceAttr(pipeline, "name", name),
					resource.TestCheckResourceAttr(pipeline, "application", app),
					testAccCheckPipelineTriggers(pipeline, []string{
						trigger,
					}),
				),
			},
			{
				Config: testAccPipelineConfigTrigger(app, newName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipeline),
					resource.TestCheckResourceAttr(pipeline, "name", newName),
					resource.TestCheckResourceAttr(pipeline, "application", app),
					testAccCheckPipelineTriggers(pipeline, []string{
						trigger,
					}),
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

  parameter {
    name = "My %s parameter"
  }

  parameter {
    name        = "Detailed parameter"
	description = "Setting options"

	default = "mosdef"
	label   = "whatevs"

	option {
	  value = 1
	}
	option {
	  value = "two"
	}
  }
}
`, app, name, name)
}

func testAccPipelineConfigTrigger(app string, name string) string {
	return fmt.Sprintf(`
resource "spinnaker_pipeline" "test" {
	application = "%s"
	name        = "%s"
	index       = 2
}

resource "spinnaker_pipeline_trigger" "jenkins" {
	pipeline = "${spinnaker_pipeline.test.id}"
	job = "Bridge Career/job/Bridge_nav/job/Bridge_nav_postmerge"
	master = "inst-ci"
	type = "jenkins"
}
`, app, name)
}

func testAccCheckPipelineExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		pipelineService := testAccProvider.Meta().(*Services).PipelineService
		_, err := pipelineService.GetPipeline(rs.Primary.Attributes["application"], rs.Primary.Attributes["name"])
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckPipelineParameters(resourceName string, expected []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		pipelineService := testAccProvider.Meta().(*Services).PipelineService
		pipeline, err := pipelineService.GetPipeline(rs.Primary.Attributes["application"], rs.Primary.Attributes["name"])
		if err != nil {
			return err
		}

		err = assertParameters(pipeline, expected)
		if err != nil {
			return err
		}

		return nil
	}
}

func assertParameters(pipeline *client.Pipeline, expected []string) error {
	if len(expected) > 0 {
		if pipeline.ParameterConfig == nil {
			return fmt.Errorf("pipeline.ParameterConfig is nil, expected: %v for pipeline:\n%+v", expected, pipeline)
		}

		if len(*pipeline.ParameterConfig) != len(expected) {
			return fmt.Errorf("pipeline.ParameterConfig is smaller than: %v", expected)
		}
	} else {
		if (pipeline.ParameterConfig != nil) && len(*pipeline.ParameterConfig) > 0 {
			return fmt.Errorf("pipeline.ParameterConfig should be empty")
		}
	}

	for _, p := range *pipeline.ParameterConfig {
		for _, s := range expected {
			if p.Name == s {
				return nil
			}
		}
	}

	return fmt.Errorf("Parameters do not match: %v, %v", expected, pipeline.ParameterConfig)
}

func testAccCheckPipelineDestroy(s *terraform.State) error {
	pipelineService := testAccProvider.Meta().(*Services).PipelineService
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "spinnaker_pipeline" {
			_, err := pipelineService.GetPipeline(rs.Primary.Attributes["application"], rs.Primary.Attributes["name"])
			if err == nil {
				return fmt.Errorf("Pipeline still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
