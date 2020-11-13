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

func TestAccPipelineTriggerBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var triggers []client.Trigger
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	triggeringPipeline := "inst-ci"
	newTriggeringPipeline := triggeringPipeline + "-new"
	trigger1 := "spinnaker_pipeline_pipeline_trigger.t1"
	trigger2 := "spinnaker_pipeline_pipeline_trigger.t2"
	pipelineResourceName := "spinnaker_pipeline.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineTriggerConfigBasic(pipeName, triggeringPipeline, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "triggering_application", "app"),
					resource.TestCheckResourceAttr(trigger1, "triggering_pipeline", triggeringPipeline),
					resource.TestCheckResourceAttr(trigger1, "status.0", "successful"),
					resource.TestCheckResourceAttr(trigger2, "triggering_application", "app"),
					resource.TestCheckResourceAttr(trigger2, "triggering_pipeline", triggeringPipeline),
					resource.TestCheckResourceAttr(trigger2, "status.0", "successful"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{
						trigger1,
						trigger2,
					}, &triggers),
				),
			},
			{
				ResourceName:  trigger1,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import key, must be pipelineID_triggerID`),
			},
			{
				ResourceName: trigger1,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(triggers) == 0 {
						return "", fmt.Errorf("no triggers to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, triggers[0].GetID()), nil
				},
				ImportStateVerify: true,
			},
			{
				ResourceName: trigger2,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(triggers) < 2 {
						return "", fmt.Errorf("no triggers to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, triggers[1].GetID()), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccPipelineTriggerConfigBasic(pipeName, newTriggeringPipeline, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "triggering_application", "app"),
					resource.TestCheckResourceAttr(trigger1, "triggering_pipeline", newTriggeringPipeline),
					resource.TestCheckResourceAttr(trigger1, "status.0", "successful"),
					resource.TestCheckResourceAttr(trigger2, "triggering_application", "app"),
					resource.TestCheckResourceAttr(trigger2, "triggering_pipeline", newTriggeringPipeline),
					resource.TestCheckResourceAttr(trigger2, "status.0", "successful"),

					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{
						trigger1,
						trigger2,
					}, &triggers),
				),
			},
			{
				Config: testAccPipelineTriggerConfigBasic(pipeName, triggeringPipeline, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "triggering_application", "app"),
					resource.TestCheckResourceAttr(trigger1, "triggering_pipeline", triggeringPipeline),
					resource.TestCheckResourceAttr(trigger1, "status.0", "successful"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{
						trigger1,
					}, &triggers),
				),
			},
			{
				Config: testAccPipelineTriggerConfigBasic(pipeName, triggeringPipeline, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{}, &triggers),
				),
			},
		},
	})
}

func testAccPipelineTriggerConfigBasic(pipeName string, triggeringPipeline string, count int) string {
	triggers := ""
	for i := 1; i <= count; i++ {
		triggers += fmt.Sprintf(`
resource "spinnaker_pipeline_pipeline_trigger" "t%v" {
	pipeline = "${spinnaker_pipeline.test.id}"

	triggering_application = "app"
	triggering_pipeline = "%s"
	status = ["successful"]
}`, i, triggeringPipeline)
	}

	return testAccPipelineConfigBasic("app", pipeName) + triggers
}
