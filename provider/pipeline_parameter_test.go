package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func TestAccPipelineParameterBasic(t *testing.T) {
	parameter1 := "spinnaker_pipeline_parameter.1"
	parameter2 := "spinnaker_pipeline_parameter.2"
	pipeline := "spinnaker_pipeline.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineParameterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineParameterConfigBasic("name", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(parameter1, "name", "name-1"),
					resource.TestCheckResourceAttr(parameter2, "name", "name-2"),
					testAccCheckPipelineParameters(pipeline, []string{
						parameter1,
						parameter2,
					}),
				),
			},
			{
				ResourceName:      parameter1,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      parameter2,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPipelineParameterConfigBasic("new-name", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(parameter1, "name", "new-name-1"),
					resource.TestCheckResourceAttr(parameter2, "name", "new-name-2"),
					testAccCheckPipelineParameters(pipeline, []string{
						parameter1,
						parameter2,
					}),
				),
			},
			{
				Config: testAccPipelineParameterConfigBasic("name", 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(parameter1, "name", "name-1"),
					testAccCheckPipelineParameters(pipeline, []string{
						parameter1,
					}),
				),
			},
			{
				Config: testAccPipelineParameterConfigBasic("name", 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineParameters(pipeline, []string{}),
				),
			},
		},
	})
}

func testAccPipelineParameterConfigBasic(name string, count int) string {
	parameters := ""
	for i := 1; i <= count; i++ {
		parameters += fmt.Sprintf(`
resource "spinnaker_pipeline_parameter" "%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name = "%s-%s"
}`, i, name, i)
	}

	return fmt.Sprintf(`
resource "spinnaker_pipeline" "test" {
	application = "app"
	name        = "pipe"
	index       = 3
}
%s
`, parameters)
}

func testAccCheckPipelineParameters(resourceName string, expected []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Pipeline not found: %s", resourceName)
		}

		pipelineService := testAccProvider.Meta().(*Services).PipelineService
		pipeline, err := pipelineService.GetPipelineByID(rs.Primary.Attributes["id"])
		if err != nil {
			return err
		}

		if len(expected) != len(*pipeline.ParameterConfig) {
			return fmt.Errorf("Parameter count of %v is expected to be %v",
				len(*pipeline.ParameterConfig), len(expected))
		}

		for _, parameterResourceName := range expected {
			expectedResource, ok := s.RootModule().Resources[parameterResourceName]
			if !ok {
				return fmt.Errorf("Parameter not found: %s", resourceName)
			}

			err = ensureParameter(pipeline.ParameterConfig, expectedResource)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func ensureParameter(parameters *[]*client.PipelineParameter, expected *terraform.ResourceState) error {
	expectedID := expected.Primary.Attributes["id"]
	for _, parameter := range *parameters {
		if parameter.ID == expectedID {
			return nil
		}
	}
	return fmt.Errorf("Parameter not found %s", expectedID)
}

func testAccCheckPipelineParameterDestroy(s *terraform.State) error {
	pipelineService := testAccProvider.Meta().(*Services).PipelineService
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "spinnaker_pipeline_parameter" {
			_, err := pipelineService.GetPipelineByID(rs.Primary.Attributes[PipelineKey])
			if err == nil {
				return fmt.Errorf("Pipeline parameter still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
