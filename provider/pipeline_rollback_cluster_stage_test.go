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
	stageTypes["spinnaker_pipeline_rollback_cluster_stage"] = client.RollbackClusterStageType
}

func TestAccPipelineRollbackClusterStageBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var stages []client.Stage
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	targetHealthyRollbackPercentage := "95"
	newTargetHealthyRollbackPercentage := "90"
	pipelineResourceName := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_rollback_cluster_stage.s1"
	stage2 := "spinnaker_pipeline_rollback_cluster_stage.s2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineRollbackClusterStageConfigBasic(pipeName, targetHealthyRollbackPercentage, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target_healthy_rollback_percentage", targetHealthyRollbackPercentage),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "target_healthy_rollback_percentage", targetHealthyRollbackPercentage),
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
				Config: testAccPipelineRollbackClusterStageConfigBasic(pipeName, newTargetHealthyRollbackPercentage, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target_healthy_rollback_percentage", newTargetHealthyRollbackPercentage),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "target_healthy_rollback_percentage", newTargetHealthyRollbackPercentage),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
						stage2,
					}, &stages),
				),
			},
			{
				Config: testAccPipelineRollbackClusterStageConfigBasic(pipeName, targetHealthyRollbackPercentage, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "target_healthy_rollback_percentage", targetHealthyRollbackPercentage),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
					}, &stages),
				),
			},
			{
				Config: testAccPipelineRollbackClusterStageConfigBasic(pipeName, targetHealthyRollbackPercentage, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{}, &stages),
				),
			},
		},
	})
}

func testAccPipelineRollbackClusterStageConfigBasic(pipeName string, targetHealthyRollbackPercentage string, count int) string {
	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_rollback_cluster_stage" "s%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Stage %v"
	cluster  = "test_cluster"
	target_healthy_rollback_percentage = %v
}`, i, i, targetHealthyRollbackPercentage)
	}

	return testAccPipelineConfigBasic("app", pipeName) + stages
}
