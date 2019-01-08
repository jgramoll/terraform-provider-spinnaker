package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func init() {
	stageTypes["spinnaker_pipeline_rollback_cluster_stage"] = client.RollbackClusterType
}

func TestAccPipelineRollbackClusterStageBasic(t *testing.T) {
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	targetHealthyRollbackPercentage := "95"
	newTargetHealthyRollbackPercentage := "90"
	pipeline := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_rollback_cluster_stage.1"
	stage2 := "spinnaker_pipeline_rollback_cluster_stage.2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineStageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineRollbackClusterStageConfigBasic(pipeName, targetHealthyRollbackPercentage, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target_healthy_rollback_percentage", targetHealthyRollbackPercentage),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "target_healthy_rollback_percentage", targetHealthyRollbackPercentage),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
						stage2,
					}),
				),
			},
			{
				Config: testAccPipelineRollbackClusterStageConfigBasic(pipeName, newTargetHealthyRollbackPercentage, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target_healthy_rollback_percentage", newTargetHealthyRollbackPercentage),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "target_healthy_rollback_percentage", newTargetHealthyRollbackPercentage),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
						stage2,
					}),
				),
			},
			{
				Config: testAccPipelineRollbackClusterStageConfigBasic(pipeName, targetHealthyRollbackPercentage, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target_healthy_rollback_percentage", targetHealthyRollbackPercentage),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
					}),
				),
			},
			{
				Config: testAccPipelineRollbackClusterStageConfigBasic(pipeName, targetHealthyRollbackPercentage, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineStages(pipeline, []string{}),
				),
			},
		},
	})
}

func testAccPipelineRollbackClusterStageConfigBasic(pipeName string, targetHealthyRollbackPercentage string, count int) string {
	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_rollback_cluster_stage" "%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Stage %v"
	cluster  = "test_cluster"
	target_healthy_rollback_percentage = %v
}`, i, i, targetHealthyRollbackPercentage)
	}

	return fmt.Sprintf(`
resource "spinnaker_pipeline" "test" {
	application = "app"
	name        = "%s"
}`, pipeName) + stages
}
