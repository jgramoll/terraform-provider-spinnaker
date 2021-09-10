package provider

import (
	"fmt"

	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var stageTypes = map[string]client.StageType{}

func testAccCheckPipelineStages(resourceName string, expected []string, stages *[]client.Stage) resource.TestCheckFunc {
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

		if len(expected) != len(*pipeline.Stages) {
			return fmt.Errorf("Stages count of %v is expected to be %v",
				len(*pipeline.Stages), len(expected))
		}

		for _, stageResourceName := range expected {
			expectedResource, ok := s.RootModule().Resources[stageResourceName]
			if !ok {
				return fmt.Errorf("Stage not found: %s", resourceName)
			}

			stage, err := ensureStage(pipeline, expectedResource)
			if err != nil {
				return err
			}
			*stages = append(*stages, stage)
		}

		return nil
	}
}

func ensureStage(pipeline *client.Pipeline, expected *terraform.ResourceState) (client.Stage, error) {
	stage, err := pipeline.GetStage(expected.Primary.Attributes["id"])
	if err != nil {
		return nil, err
	}

	var expectedType = stageTypes[expected.Type]
	if expectedType != stage.GetType() {
		return nil, fmt.Errorf("Stage %s has expected type \"%s\", got type \"%s\"",
			stage.GetRefID(), expectedType, stage.GetType())
	}
	return stage, nil
}
