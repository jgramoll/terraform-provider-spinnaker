package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func init() {
	stageTypes["spinnaker_pipeline_pipeline_stage"] = client.PipelineStageType
}

func TestAccPipelinePipelineStageBasic(t *testing.T) {
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	targetPipeline := "my-pipeline"
	newTargetPipeline := "new-pipeline"
	pipeline := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_pipeline_stage.1"
	stage2 := "spinnaker_pipeline_pipeline_stage.2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineStageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelinePipelineStageConfigBasic(pipeName, targetPipeline, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target_pipeline", targetPipeline),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "target_pipeline", targetPipeline),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
						stage2,
					}),
				),
			},
			{
				ResourceName:      stage1,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      stage2,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPipelinePipelineStageConfigBasic(pipeName, newTargetPipeline, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target_pipeline", newTargetPipeline),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "target_pipeline", newTargetPipeline),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
						stage2,
					}),
				),
			},
			{
				Config: testAccPipelinePipelineStageConfigBasic(pipeName, targetPipeline, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target_pipeline", targetPipeline),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
					}),
				),
			},
			{
				Config: testAccPipelinePipelineStageConfigBasic(pipeName, targetPipeline, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineStages(pipeline, []string{}),
				),
			},
			{
				ResourceName:      pipeline,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPipelinePipelineStageConfigBasic(pipeName string, targetPipeline string, count int) string {
	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_pipeline_stage" "%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Stage %v"
	application  = "test_app"
	target_pipeline = "%v"
}`, i, i, targetPipeline)
	}

	return fmt.Sprintf(`
resource "spinnaker_pipeline" "test" {
	application = "app"
	name        = "%s"
}`, pipeName) + stages
}
