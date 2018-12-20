package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func TestAccPipelineStageBasic(t *testing.T) {
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	vmType := "hvm"
	newVMType := "pv"
	pipeline := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_bake_stage.1"
	stage2 := "spinnaker_pipeline_bake_stage.2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineStageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineStageConfigBasic(pipeName, vmType, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "vm_type", vmType),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "vm_type", vmType),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
						stage2,
					}),
				),
			},
			{
				Config: testAccPipelineStageConfigBasic(pipeName, newVMType, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "vm_type", newVMType),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "vm_type", newVMType),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
						stage2,
					}),
				),
			},
			{
				Config: testAccPipelineStageConfigBasic(pipeName, vmType, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "vm_type", vmType),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
					}),
				),
			},
			{
				Config: testAccPipelineStageConfigBasic(pipeName, vmType, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineStages(pipeline, []string{}),
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

var stageTypes = map[string]client.StageType{
	"spinnaker_pipeline_bake_stage": client.BakeStageType,
}

func ensureStage(stages []client.Stage, expected *terraform.ResourceState) error {
	expectedID := expected.Primary.Attributes["id"]
	for _, s := range stages {
		if s.GetRefID() == expectedID {
			var expectedType = stageTypes[expected.Type]
			if expectedType != s.GetType() {
				return fmt.Errorf("Stage %s has expected type %s, got type \"%s\"",
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
		if rs.Type == "spinnaker_pipeline_bake_stage" {
			_, err := pipelineService.GetPipelineByID(rs.Primary.Attributes[PipelineKey])
			if err == nil {
				return fmt.Errorf("Pipeline stage still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
