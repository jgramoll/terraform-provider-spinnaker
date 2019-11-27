package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func init() {
	stageTypes["spinnaker_pipeline_check_preconditions_stage"] = client.CheckPreconditionsStageType
}

func TestAccPipelineCheckPreconditionsStageBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var stages []client.Stage
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	stageName := "my-stage"
	// newAccountName := accountName + "-new"
	pipelineResourceName := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_check_preconditions_stage.s1"
	stage2 := "spinnaker_pipeline_check_preconditions_stage.s2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineCheckPreconditionsStageConfigBasic(pipeName, stageName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "precondition.0.type", "stageStatus"),
					resource.TestCheckResourceAttr(stage1, "precondition.0.context.stage_name", stageName),
					resource.TestCheckResourceAttr(stage1, "precondition.0.context.stage_status", "FAILED_CONTINUE"),

					// precondition {
					// 	context {
					// 		stage_name   = "Manual Judgment"
					// 		stage_status = "FAILED_CONTINUE"
					// 	}
					// 	type = "stageStatus"
					// }
					// precondition {
					// 	context {
					// 		expression = "this is myexp"
					// 	}
					// 	type = "expression"
					// }
					// precondition {
					// 	context {
					// 		expected = 1
					// 	}
					// 	type = "clusterSize"
					// }

					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
						stage2,
					}, &stages),
				),
			},
			// {
			// 	ResourceName:  stage1,
			// 	ImportStateId: "invalid",
			// 	ImportState:   true,
			// 	ExpectError:   regexp.MustCompile(`Invalid import key, must be pipelineID_stageID`),
			// },
			// {
			// 	ResourceName: stage1,
			// 	ImportState:  true,
			// 	ImportStateIdFunc: func(*terraform.State) (string, error) {
			// 		if len(stages) == 0 {
			// 			return "", fmt.Errorf("no stages to import")
			// 		}
			// 		return fmt.Sprintf("%s_%s", pipelineRef.ID, stages[0].GetRefID()), nil
			// 	},
			// 	ImportStateVerify: true,
			// },
			// {
			// 	ResourceName: stage2,
			// 	ImportState:  true,
			// 	ImportStateIdFunc: func(*terraform.State) (string, error) {
			// 		if len(stages) < 2 {
			// 			return "", fmt.Errorf("no stages to import")
			// 		}
			// 		return fmt.Sprintf("%s_%s", pipelineRef.ID, stages[1].GetRefID()), nil
			// 	},
			// 	ImportStateVerify: true,
			// },
			// {
			// 	Config: testAccPipelineCheckPreconditionsStageConfigBasic(pipeName, newAccountName, 2),
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
			// 		resource.TestCheckResourceAttr(stage1, "account", newAccountName),
			// 		resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
			// 		resource.TestCheckResourceAttr(stage2, "account", newAccountName),
			// 		testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
			// 		testAccCheckPipelineStages(pipelineResourceName, []string{
			// 			stage1,
			// 			stage2,
			// 		}, &stages),
			// 	),
			// },
			// {
			// 	Config: testAccPipelineCheckPreconditionsStageConfigBasic(pipeName, accountName, 1),
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
			// 		resource.TestCheckResourceAttr(stage1, "account", accountName),
			// 		testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
			// 		testAccCheckPipelineStages(pipelineResourceName, []string{
			// 			stage1,
			// 		}, &stages),
			// 	),
			// },
			// {
			// 	Config: testAccPipelineCheckPreconditionsStageConfigBasic(pipeName, accountName, 0),
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
			// 		testAccCheckPipelineStages(pipelineResourceName, []string{}, &stages),
			// 	),
			// },
		},
	})
}

func testAccPipelineCheckPreconditionsStageConfigBasic(pipeName string, stageName string, count int) string {
	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_check_preconditions_stage" "s%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Stage %v"
	
	precondition {
		type = "stageStatus"

		context = {
			stage_name   = "%s"
			stage_status = "FAILED_CONTINUE"
		}
	}
	precondition {
		type = "expression"

		context = {
			expression = "this is myexp"
		}
	}
	precondition {
		type = "clusterSize"

		context = {
			credentials = "my-cred"
			expected = 1
			regions = "us-east-1, us-east-2"
		}
	}
}`, i, i, stageName)
	}

	return fmt.Sprintf(`
resource "spinnaker_pipeline" "test" {
	application = "app"
	name        = "%s"
}`, pipeName) + stages
}
