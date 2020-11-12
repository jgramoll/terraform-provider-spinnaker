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

func TestAccDockerTriggerBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var triggers []client.Trigger
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	account := "my-account"
	newAccount := account + "-new"
	trigger1 := "spinnaker_pipeline_docker_trigger.t1"
	trigger2 := "spinnaker_pipeline_docker_trigger.t2"
	pipelineResourceName := "spinnaker_pipeline.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDockerTriggerConfigBasic(pipeName, account, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "account", account),
					resource.TestCheckResourceAttr(trigger1, "organization", "my-org"),
					resource.TestCheckResourceAttr(trigger1, "repository", "repository"),
					resource.TestCheckResourceAttr(trigger2, "account", account),
					resource.TestCheckResourceAttr(trigger2, "organization", "my-org"),
					resource.TestCheckResourceAttr(trigger2, "repository", "repository"),
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
				Config: testAccDockerTriggerConfigBasic(pipeName, newAccount, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "account", newAccount),
					resource.TestCheckResourceAttr(trigger1, "organization", "my-org"),
					resource.TestCheckResourceAttr(trigger1, "repository", "repository"),
					resource.TestCheckResourceAttr(trigger2, "account", newAccount),
					resource.TestCheckResourceAttr(trigger2, "organization", "my-org"),
					resource.TestCheckResourceAttr(trigger2, "repository", "repository"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{
						trigger1,
						trigger2,
					}, &triggers),
				),
			},
			{
				Config: testAccDockerTriggerConfigBasic(pipeName, account, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "account", account),
					resource.TestCheckResourceAttr(trigger1, "organization", "my-org"),
					resource.TestCheckResourceAttr(trigger1, "repository", "repository"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{
						trigger1,
					}, &triggers),
				),
			},
			{
				Config: testAccDockerTriggerConfigBasic(pipeName, account, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{}, &triggers),
				),
			},
		},
	})
}

func testAccDockerTriggerConfigBasic(pipeName string, master string, count int) string {
	triggers := ""
	for i := 1; i <= count; i++ {
		triggers += fmt.Sprintf(`
resource "spinnaker_pipeline_docker_trigger" "t%v" {
	pipeline = "${spinnaker_pipeline.test.id}"

	account = "%s"
	organization = "my-org"
	repository = "repository"
}`, i, master)
	}

	return testAccPipelineConfigBasic("app", pipeName) + triggers
}
