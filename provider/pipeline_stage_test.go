package provider

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

var stageTypes = map[string]client.StageType{}

func testAccCheckPipelineStages(resourceName string, expected []string) resource.TestCheckFunc {
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

		if len(expected) != len(pipeline.Stages) {
			return fmt.Errorf("Stages count of %v is expected to be %v",
				len(pipeline.Stages), len(expected))
		}

		for _, stageResourceName := range expected {
			expectedResource, ok := s.RootModule().Resources[stageResourceName]
			if !ok {
				return fmt.Errorf("Stage not found: %s", resourceName)
			}

			err = ensureStage(pipeline.Stages, expectedResource)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func ensureStage(stages []client.Stage, expected *terraform.ResourceState) error {
	expectedID := expected.Primary.Attributes["id"]
	for _, s := range stages {
		if s.GetRefID() == expectedID {
			var expectedType = stageTypes[expected.Type]
			if expectedType != s.GetType() {
				return fmt.Errorf("Stage %s has expected type \"%s\", got type \"%s\"",
					s.GetRefID(), expectedType, s.GetType())
			}
			return nil
		}
	}
	return fmt.Errorf("Stage not found %s", expectedID)
}

func testAccCheckPipelineStageDestroy(s *terraform.State) error {
	pipelineService := testAccProvider.Meta().(*Services).PipelineService
	for _, rs := range s.RootModule().Resources {
		if _, ok := stageTypes[rs.Type]; ok {
			_, err := pipelineService.GetPipelineByID(rs.Primary.Attributes[PipelineKey])
			if err == nil {
				return fmt.Errorf("Pipeline stage still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
