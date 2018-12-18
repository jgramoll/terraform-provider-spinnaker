package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func TestAccPipelineStage_basic(t *testing.T) {
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineStageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineStageConfigBasic(pipeName, "hvm", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("spinnaker_pipeline_bake_stage.1", "name", "Stage 1"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_bake_stage.1", "vm_type", "hvm"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_bake_stage.2", "name", "Stage 2"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_bake_stage.2", "vm_type", "hvm"),
					testAccCheckPipelineStages("spinnaker_pipeline.test", []string{
						"spinnaker_pipeline_bake_stage.1",
						"spinnaker_pipeline_bake_stage.2",
					}),
				),
			},
			{
				Config: testAccPipelineStageConfigBasic(pipeName, "pv", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("spinnaker_pipeline_bake_stage.1", "name", "Stage 1"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_bake_stage.1", "vm_type", "pv"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_bake_stage.2", "name", "Stage 2"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_bake_stage.2", "vm_type", "pv"),
					testAccCheckPipelineStages("spinnaker_pipeline.test", []string{
						"spinnaker_pipeline_bake_stage.1",
						"spinnaker_pipeline_bake_stage.2",
					}),
				),
			},
			{
				Config: testAccPipelineStageConfigBasic(pipeName, "hvm", 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("spinnaker_pipeline_bake_stage.1", "name", "Stage 1"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_bake_stage.1", "vm_type", "hvm"),
					testAccCheckPipelineStages("spinnaker_pipeline.test", []string{
						"spinnaker_pipeline_bake_stage.1",
					}),
				),
			},
			{
				Config: testAccPipelineStageConfigBasic(pipeName, "hvm", 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineStages("spinnaker_pipeline.test", []string{}),
				),
			},
		},
	})
}

func testAccPipelineStageConfigBasic(pipeName string, vmType string, count int) string {
	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_bake_stage" "%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Stage %v"
	vm_type  = "%v"
}`, i, i, vmType)
	}

	return fmt.Sprintf(`
resource "spinnaker_pipeline" "test" {
	application = "app"
	name        = "%s"
}`, pipeName) + stages
}

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

var stageTypes = map[string]string{
	"spinnaker_pipeline_bake_stage": "bake",
}

func ensureStage(stages []client.Stage, expected *terraform.ResourceState) error {
	expectedID := expected.Primary.Attributes["id"]
	// for _, stage := range stages {
	// if stage.RefID == expectedID {
	// 	var expectedType = stageTypes[expected.Type]
	// 	if expectedType != stage.Type {
	// 		return fmt.Errorf("Stage %s has expected type %s, got type \"%s\"",
	// 			stage.RefID, expectedType, stage.Type)
	// 	}
	// 	return nil
	// }
	// }
	return fmt.Errorf("Stage not found %s", expectedID)
}

func testAccCheckPipelineStageDestroy(s *terraform.State) error {
	pipelineService := testAccProvider.Meta().(*Services).PipelineService
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "spinnaker_pipeline_bake_stage" {
			_, err := pipelineService.GetPipelineByID(rs.Primary.Attributes["pipeline"])
			if err == nil {
				return fmt.Errorf("Pipeline stage still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
