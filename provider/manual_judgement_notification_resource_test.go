package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func init() {
	stageTypes["spinnaker_pipeline_manual_judgment_stage"] = client.ManualJudgmentStageType
}

func TestAccPipelineManualJudgementNotificationBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var stages []client.Stage
	pipelineName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	pipelineResourceName := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_manual_judgment_stage.s1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineManualJudgmentNotificationConfigBasic(pipelineName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "notification.0.message.0.manual_judgment_continue", "Manual judgement continue"),
					resource.TestCheckResourceAttr(stage1, "notification.0.when.0.manual_judgment_continue", "true"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineStages(pipelineResourceName, []string{
						stage1,
					}, &stages),
				),
			},
		},
	})
}

func testAccPipelineManualJudgmentNotificationConfigBasic(pipelineName string) string {
	var stages = `
resource "spinnaker_pipeline_manual_judgment_stage" "s1" {
	pipeline 	 = "${spinnaker_pipeline.test.id}"
	name     	 = "Stage 1"
	instructions = "Manual Judgment Instructions"

	judgment_inputs = [
		"commit",
		"rollback",
	]

	notification {
		address = "#my-slack-channel"
		message {
			manual_judgment_continue = "Manual judgement continue"
			manual_judgment_stop = "Manual judgement stop"
		}
		type = "slack"
		when {
			manual_judgment = true
			manual_judgment_continue = true
			manual_judgment_stop = true
		}
	}

	notification {
		address = "#my-other-channel"
		message {}
		type = "slack"
		when {
			manual_judgment = false
		}
	}
}`

	return testAccPipelineConfigBasic("app", pipelineName) + stages
}
