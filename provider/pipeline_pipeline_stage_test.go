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

func init() {
	stageTypes["spinnaker_pipeline_pipeline_stage"] = client.PipelineStageType
}

func TestAccPipelinePipelineStageBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var stages []client.Stage
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	targetPipeline := "my-pipeline"
	newTargetPipeline := "new-pipeline"
	pipelineResourceName := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_pipeline_stage.s1"
	stage2 := "spinnaker_pipeline_pipeline_stage.s2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelinePipelineStageConfigBasic(pipeName, targetPipeline, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target_pipeline", targetPipeline),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "target_pipeline", targetPipeline),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
						stage2,
					}, &stages),
				),
			},
			{
				ResourceName:  stage1,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import key, must be pipelineID_stageID`),
			},
			{
				ResourceName: stage1,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(stages) == 0 {
						return "", fmt.Errorf("no stages to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, stages[0].GetRefID()), nil
				},
				ImportStateVerify: true,
			},
			{
				ResourceName: stage2,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(stages) < 2 {
						return "", fmt.Errorf("no stages to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, stages[1].GetRefID()), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccPipelinePipelineStageConfigBasic(pipeName, newTargetPipeline, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target_pipeline", newTargetPipeline),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "target_pipeline", newTargetPipeline),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
						stage2,
					}, &stages),
				),
			},
			{
				Config: testAccPipelinePipelineStageConfigBasic(pipeName, targetPipeline, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target_pipeline", targetPipeline),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
					}, &stages),
				),
			},
			{
				Config: testAccPipelinePipelineStageConfigBasic(pipeName, targetPipeline, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{}, &stages),
				),
			},
			{
				ResourceName:      pipelineResourceName,
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
resource "spinnaker_pipeline_pipeline_stage" "s%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Stage %v"
	application  = "test_app"
	target_pipeline = "%v"
}`, i, i, targetPipeline)
	}

	return testAccPipelineConfigBasic("app", pipeName) + stages
}
