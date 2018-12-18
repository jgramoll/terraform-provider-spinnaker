package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func TestAccPipelineTrigger_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineTriggerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineTriggerConfigBasic("inst-ci", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("spinnaker_pipeline_trigger.1", "master", "inst-ci"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_trigger.1", "property_file", "build.properties.1"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_trigger.2", "master", "inst-ci"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_trigger.2", "property_file", "build.properties.2"),
					testAccCheckPipelineTriggers("spinnaker_pipeline.test", []string{
						"spinnaker_pipeline_trigger.1",
						"spinnaker_pipeline_trigger.2",
					}),
				),
			},
			{
				Config: testAccPipelineTriggerConfigBasic("inst-ci-new", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("spinnaker_pipeline_trigger.1", "master", "inst-ci-new"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_trigger.1", "property_file", "build.properties.1"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_trigger.2", "master", "inst-ci-new"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_trigger.2", "property_file", "build.properties.2"),
					testAccCheckPipelineTriggers("spinnaker_pipeline.test", []string{
						"spinnaker_pipeline_trigger.1",
						"spinnaker_pipeline_trigger.2",
					}),
				),
			},
			{
				Config: testAccPipelineTriggerConfigBasic("inst-ci", 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("spinnaker_pipeline_trigger.1", "master", "inst-ci"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_trigger.1", "property_file", "build.properties.1"),
					testAccCheckPipelineTriggers("spinnaker_pipeline.test", []string{
						"spinnaker_pipeline_trigger.1",
					}),
				),
			},
			{
				Config: testAccPipelineTriggerConfigBasic("inst-ci", 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineTriggers("spinnaker_pipeline.test", []string{}),
				),
			},
		},
	})
}

func testAccPipelineTriggerConfigBasic(master string, count int) string {
	triggers := ""
	for i := 1; i <= count; i++ {
		triggers += fmt.Sprintf(`
resource "spinnaker_pipeline_trigger" "%v" {
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
	name        = "pipe"
	index       = 3
}
%s
`, triggers)
}

func testAccCheckPipelineTriggers(resourceName string, expected []string) resource.TestCheckFunc {
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

			err = ensureTrigger(pipeline.Triggers, expectedResource)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func ensureTrigger(triggers []client.Trigger, expected *terraform.ResourceState) error {
	expectedID := expected.Primary.Attributes["id"]
	for _, trigger := range triggers {
		if trigger.ID == expectedID {
			return nil
		}
	}
	return fmt.Errorf("Trigger not found %s", expectedID)
}

func testAccCheckPipelineTriggerDestroy(s *terraform.State) error {
	pipelineService := testAccProvider.Meta().(*Services).PipelineService
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "spinnaker_pipeline_trigger" {
			_, err := pipelineService.GetPipelineByID(rs.Primary.Attributes["pipeline"])
			if err == nil {
				return fmt.Errorf("Pipeline trigger still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
