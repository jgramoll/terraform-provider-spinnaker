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

func TestAccWebhookTriggerBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var triggers []client.Trigger
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	source := "my-source"
	newSource := source + "-new"
	trigger1 := "spinnaker_pipeline_webhook_trigger.t1"
	trigger2 := "spinnaker_pipeline_webhook_trigger.t2"
	pipelineResourceName := "spinnaker_pipeline.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccWebhookTriggerConfigBasic(pipeName, source, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "source", source),
					resource.TestCheckResourceAttr(trigger1, "payload_constraints.foo", "bar"),
					resource.TestCheckResourceAttr(trigger1, "payload_constraints.baz", "qux"),
					resource.TestCheckResourceAttr(trigger2, "source", source),
					resource.TestCheckResourceAttr(trigger2, "payload_constraints.foo", "bar"),
					resource.TestCheckResourceAttr(trigger2, "payload_constraints.baz", "qux"),

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
				Config: testAccWebhookTriggerConfigBasic(pipeName, newSource, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "source", newSource),
					resource.TestCheckResourceAttr(trigger1, "payload_constraints.foo", "bar"),
					resource.TestCheckResourceAttr(trigger1, "payload_constraints.baz", "qux"),
					resource.TestCheckResourceAttr(trigger2, "source", newSource),
					resource.TestCheckResourceAttr(trigger2, "payload_constraints.foo", "bar"),
					resource.TestCheckResourceAttr(trigger2, "payload_constraints.baz", "qux"),

					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{
						trigger1,
						trigger2,
					}, &triggers),
				),
			},
			{
				Config: testAccWebhookTriggerConfigBasic(pipeName, source, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "source", source),
					resource.TestCheckResourceAttr(trigger1, "payload_constraints.foo", "bar"),
					resource.TestCheckResourceAttr(trigger1, "payload_constraints.baz", "qux"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{
						trigger1,
					}, &triggers),
				),
			},
			{
				Config: testAccWebhookTriggerConfigBasic(pipeName, source, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{}, &triggers),
				),
			},
		},
	})
}

func testAccWebhookTriggerConfigBasic(pipeName string, source string, count int) string {
	triggers := ""
	for i := 1; i <= count; i++ {
		triggers += fmt.Sprintf(`
resource "spinnaker_pipeline_webhook_trigger" "t%v" {
	pipeline = spinnaker_pipeline.test.id

	source = "%s"
	
	payload_constraints = {
		"foo" = "bar"
		"baz" = "qux"
	}
}`, i, source)
	}

	return testAccPipelineConfigBasic("app", pipeName) + triggers
}
