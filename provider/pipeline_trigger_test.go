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
	var triggers []*client.Trigger
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	jenkinsMaster := "inst-ci"
	newJenkinsMaster := jenkinsMaster + "-new"
	trigger1 := "spinnaker_pipeline_trigger.trigger-1"
	trigger2 := "spinnaker_pipeline_trigger.trigger-2"
	pipelineResourceName := "spinnaker_pipeline.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineTriggerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineTriggerConfigBasic(pipeName, jenkinsMaster, 2),
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
					return fmt.Sprintf("%s_%s", pipelineRef.ID, triggers[0].ID), nil
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
					return fmt.Sprintf("%s_%s", pipelineRef.ID, triggers[1].ID), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccPipelineTriggerConfigBasic(pipeName, newJenkinsMaster, 2),
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
				Config: testAccPipelineTriggerConfigBasic(pipeName, jenkinsMaster, 1),
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
				Config: testAccPipelineTriggerConfigBasic(pipeName, jenkinsMaster, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineTriggers(pipelineResourceName, []string{}, &triggers),
				),
			},
		},
	})
}

func testAccPipelineTriggerConfigBasic(pipeName string, master string, count int) string {
	triggers := ""
	for i := 1; i <= count; i++ {
		triggers += fmt.Sprintf(`
resource "spinnaker_pipeline_trigger" "trigger-%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	job = "Bridge Career/job/Bridge_nav/job/Bridge_nav_postmerge"
	master = "%s"
	property_file = "build.properties.%v"
	type = "jenkins"
}`, i, master, i)
	}

	return fmt.Sprintf(`
resource "spinnaker_pipeline" "test" {
	application = "app"
	name        = "%s"
}`, pipeName) + triggers
}

func testAccCheckPipelineTriggers(resourceName string, expected []string, triggers *[]*client.Trigger) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Pipeline not found: %s", resourceName)
		}

		pipelineService := testAccProvider.Meta().(*Services).PipelineService
		pipeline, err := pipelineService.GetPipelineByID(rs.Primary.Attributes["id"])
		if err != nil {
			return err
		}

		if len(expected) != len(pipeline.Triggers) {
			return fmt.Errorf("Trigger count of %v is expected to be %v",
				len(pipeline.Triggers), len(expected))
		}

		for _, triggerResourceName := range expected {
			expectedResource, ok := s.RootModule().Resources[triggerResourceName]
			if !ok {
				return fmt.Errorf("Trigger not found: %s", resourceName)
			}

			t, err := ensureTrigger(&pipeline.Triggers, expectedResource)
			if err != nil {
				return err
			}
			*triggers = append(*triggers, t)
		}

		return nil
	}
}

func ensureTrigger(triggers *[]*client.Trigger, expected *terraform.ResourceState) (*client.Trigger, error) {
	expectedID := expected.Primary.Attributes["id"]
	for _, t := range *triggers {
		if t.ID == expectedID {
			return t, nil
		}
	}
	return nil, fmt.Errorf("Trigger not found %s", expectedID)
}

func testAccCheckPipelineTriggerDestroy(s *terraform.State) error {
	pipelineService := testAccProvider.Meta().(*Services).PipelineService
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "spinnaker_pipeline_trigger" {
			_, err := pipelineService.GetPipelineByID(rs.Primary.Attributes[PipelineKey])
			if err == nil {
				return fmt.Errorf("Pipeline trigger still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
