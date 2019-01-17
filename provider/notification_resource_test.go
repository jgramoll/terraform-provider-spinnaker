package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccPipelineNotificationStageBasic(t *testing.T) {
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	completeText := "complete!"
	newCompleteText := completeText + "-new"
	pipeline := "spinnaker_pipeline.test"
	stage1 := "spinnaker_pipeline_jenkins_stage.1"
	stage2 := "spinnaker_pipeline_jenkins_stage.2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineStageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineNotificationStageConfigBasic(pipeName, completeText, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "notification.0.message.0.complete", completeText),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "notification.0.message.0.complete", completeText),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
						stage2,
					}),
				),
			},
			{
				Config: testAccPipelineNotificationStageConfigBasic(pipeName, newCompleteText, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "notification.0.message.0.complete", newCompleteText),
					resource.TestCheckResourceAttr(stage2, "name", "Stage 2"),
					resource.TestCheckResourceAttr(stage2, "notification.0.message.0.complete", newCompleteText),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
						stage2,
					}),
				),
			},
			{
				Config: testAccPipelineNotificationStageConfigBasic(pipeName, completeText, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(stage1, "name", "Stage 1"),
					resource.TestCheckResourceAttr(stage1, "notification.0.message.0.complete", completeText),
					testAccCheckPipelineStages(pipeline, []string{
						stage1,
					}),
				),
			},
			{
				Config: testAccPipelineNotificationStageConfigBasic(pipeName, completeText, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineStages(pipeline, []string{}),
				),
			},
		},
	})
}

func testAccPipelineNotificationStageConfigBasic(pipeName string, completeText string, count int) string {
	stages := ""
	for i := 1; i <= count; i++ {
		stages += fmt.Sprintf(`
resource "spinnaker_pipeline_jenkins_stage" "%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	name     = "Stage %v"
	master   = "master"
	job      = "jenkins/job"

	notification {
		address = "#slack-channel"
		message = {
			complete = "%v"
		}
		type = "slack"
		when = {
			complete = true
		}
	}
}`, i, i, completeText)
	}

	return fmt.Sprintf(`
resource "spinnaker_pipeline" "test" {
	application = "app"
	name        = "%s"
}`, pipeName) + stages
}
