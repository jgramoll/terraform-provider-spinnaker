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

func TestAccPipelineParameterBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var parameters []*client.PipelineParameter
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	parameter1 := "spinnaker_pipeline_parameter.p1"
	parameter2 := "spinnaker_pipeline_parameter.p2"
	pipelineResourceName := "spinnaker_pipeline.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineParameterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineParameterConfigBasic(pipeName, "name", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(parameter1, "name", "name-1"),
					resource.TestCheckResourceAttr(parameter2, "name", "name-2"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineParameters(pipelineResourceName, []string{
						parameter1,
						parameter2,
					}, &parameters),
				),
			},
			{
				ResourceName:  parameter1,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import key, must be pipelineID_parameterID`),
			},
			{
				ResourceName: parameter1,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(parameters) == 0 {
						return "", fmt.Errorf("no parameters to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, parameters[0].ID), nil
				},
				ImportStateVerify: true,
			},
			{
				ResourceName: parameter2,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(parameters) < 2 {
						return "", fmt.Errorf("no parameters to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, parameters[1].ID), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccPipelineParameterConfigBasic(pipeName, "new-name", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(parameter1, "name", "new-name-1"),
					resource.TestCheckResourceAttr(parameter2, "name", "new-name-2"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineParameters(pipelineResourceName, []string{
						parameter1,
						parameter2,
					}, &parameters),
				),
			},
			{
				Config: testAccPipelineParameterConfigBasic(pipeName, "name", 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(parameter1, "name", "name-1"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineParameters(pipelineResourceName, []string{
						parameter1,
					}, &parameters),
				),
			},
			{
				Config: testAccPipelineParameterConfigBasic(pipeName, "name", 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineParameters(pipelineResourceName, []string{}, &parameters),
				),
			},
		},
	})
}

func testAccPipelineParameterConfigBasic(pipeName string, name string, count int) string {
	parameters := ""
	for i := 1; i <= count; i++ {
		parameters += fmt.Sprintf(`
resource "spinnaker_pipeline_parameter" "p%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name = "%s-%v"
	description = "Setting options"
	default = "mosdef"
	label   = "whatevs"

	option {
	  value = 1
	}
	option {
	  value = "two"
	}
}`, i, name, i)
	}

	return testAccPipelineConfigBasic("app", pipeName) + parameters
}

func testAccCheckPipelineParameters(resourceName string, expected []string, parameters *[]*client.PipelineParameter) resource.TestCheckFunc {
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
				return fmt.Errorf("Parameter not found in resources: %s", resourceName)
			}

			parameter, err := ensureParameter(pipeline.ParameterConfig, expectedResource)
			if err != nil {
				return err
			}
			*parameters = append(*parameters, parameter)
		}

		return nil
	}
}

func ensureParameter(parameters *[]*client.PipelineParameter, expected *terraform.ResourceState) (*client.PipelineParameter, error) {
	expectedID := expected.Primary.Attributes["id"]
	for _, parameter := range *parameters {
		if parameter.ID == expectedID {
			return parameter, nil
		}
	}
	return nil, fmt.Errorf("Parameter not found in pipeline %s", expectedID)
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
