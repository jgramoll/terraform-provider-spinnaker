package provider

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func testAccCheckPipelineTriggers(resourceName string, expected []string, triggers *[]client.Trigger) resource.TestCheckFunc {
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

		if len(expected) != len(pipeline.Triggers) {
			return fmt.Errorf("Trigger count of %v is expected to be %v",
				len(pipeline.Triggers), len(expected))
		}

		for _, triggerResourceName := range expected {
			expectedResource, ok := s.RootModule().Resources[triggerResourceName]
			if !ok {
				return fmt.Errorf("Trigger not found: %s", resourceName)
			}

			t, err := ensureTrigger(&pipeline.Triggers, expectedResource)
			if err != nil {
				return err
			}
			*triggers = append(*triggers, t)
		}

		return nil
	}
}

func ensureTrigger(triggers *[]client.Trigger, expected *terraform.ResourceState) (client.Trigger, error) {
	expectedID := expected.Primary.Attributes["id"]
	for _, t := range *triggers {
		if t.GetID() == expectedID {
			return t, nil
		}
	}
	return nil, fmt.Errorf("Trigger not found %s", expectedID)
}
