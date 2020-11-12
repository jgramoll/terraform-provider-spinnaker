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

func TestAccJenkinsTriggerBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var triggers []client.Trigger
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jenkinsMaster := "inst-ci"
	newJenkinsMaster := jenkinsMaster + "-new"
	trigger1 := "spinnaker_pipeline_jenkins_trigger.t1"
	trigger2 := "spinnaker_pipeline_jenkins_trigger.t2"
	pipelineResourceName := "spinnaker_pipeline.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccJenkinsTriggerConfigBasic(pipeName, jenkinsMaster, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "master", jenkinsMaster),
					resource.TestCheckResourceAttr(trigger1, "property_file", "build.properties.1"),
					resource.TestCheckResourceAttr(trigger2, "master", jenkinsMaster),
					resource.TestCheckResourceAttr(trigger2, "property_file", "build.properties.2"),
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
				Config: testAccJenkinsTriggerConfigBasic(pipeName, newJenkinsMaster, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "master", newJenkinsMaster),
					resource.TestCheckResourceAttr(trigger1, "property_file", "build.properties.1"),
					resource.TestCheckResourceAttr(trigger2, "master", newJenkinsMaster),
					resource.TestCheckResourceAttr(trigger2, "property_file", "build.properties.2"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{
						trigger1,
						trigger2,
					}, &triggers),
				),
			},
			{
				Config: testAccJenkinsTriggerConfigBasic(pipeName, jenkinsMaster, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(trigger1, "master", jenkinsMaster),
					resource.TestCheckResourceAttr(trigger1, "property_file", "build.properties.1"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{
						trigger1,
					}, &triggers),
				),
			},
			{
				Config: testAccJenkinsTriggerConfigBasic(pipeName, jenkinsMaster, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{}, &triggers),
				),
			},
		},
	})
}

func testAccJenkinsTriggerConfigBasic(pipeName string, master string, count int) string {
	triggers := ""
	for i := 1; i <= count; i++ {
		triggers += fmt.Sprintf(`
resource "spinnaker_pipeline_jenkins_trigger" "t%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	job = "Bridge Career/job/Bridge_nav/job/Bridge_nav_postmerge"
	master = "%s"
	property_file = "build.properties.%v"
}`, i, master, i)
	}

	return testAccPipelineConfigBasic("app", pipeName) + triggers
}
