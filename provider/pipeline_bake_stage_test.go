package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccPipelineBakeStageBasic(t *testing.T) {
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
				Config: testAccPipelineBakeStageConfigBasic(pipeName, vmType, 2),
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
				Config: testAccPipelineBakeStageConfigBasic(pipeName, newVMType, 2),
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
				Config: testAccPipelineBakeStageConfigBasic(pipeName, vmType, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "vm_type", vmType),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
					}),
				),
			},
			{
				Config: testAccPipelineBakeStageConfigBasic(pipeName, vmType, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineStages(pipeline, []string{}),
				),
			},
		},
	})
}

func testAccPipelineBakeStageConfigBasic(pipeName string, vmType string, count int) string {
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
