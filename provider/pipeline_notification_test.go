package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

func TestAccPipelineNotification_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineNotificationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineNotificationConfigBasic("bridge-career-deploys", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("spinnaker_pipeline_notification.1", "address", "bridge-career-deploys"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_notification.1", "message.complete", "1 is done"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_notification.2", "address", "bridge-career-deploys"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_notification.2", "message.complete", "2 is done"),
					testAccCheckPipelineNotifications("spinnaker_pipeline.test", []string{
						"spinnaker_pipeline_notification.1",
						"spinnaker_pipeline_notification.2",
					}),
				),
			},
			{
				Config: testAccPipelineNotificationConfigBasic("bridge-career-deploys-new", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("spinnaker_pipeline_notification.1", "address", "bridge-career-deploys-new"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_notification.1", "message.complete", "1 is done"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_notification.2", "address", "bridge-career-deploys-new"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_notification.2", "message.complete", "2 is done"),
					testAccCheckPipelineNotifications("spinnaker_pipeline.test", []string{
						"spinnaker_pipeline_notification.1",
						"spinnaker_pipeline_notification.2",
					}),
				),
			},
			{
				Config: testAccPipelineNotificationConfigBasic("bridge-career-deploys", 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("spinnaker_pipeline_notification.1", "address", "bridge-career-deploys"),
					resource.TestCheckResourceAttr("spinnaker_pipeline_notification.1", "message.complete", "1 is done"),
					testAccCheckPipelineNotifications("spinnaker_pipeline.test", []string{
						"spinnaker_pipeline_notification.1",
					}),
				),
			},
			{
				Config: testAccPipelineNotificationConfigBasic("bridge-career-deploys", 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineTriggers("spinnaker_pipeline.test", []string{}),
				),
			},
		},
	})
}

func testAccPipelineNotificationConfigBasic(address string, count int) string {
	notifications := ""
	for i := 1; i <= count; i++ {
		notifications += fmt.Sprintf(`
resource "spinnaker_pipeline_notification" "%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	address = "%v"
	level = "pipeline"
	message = {
		complete = "%v is done"
		failed = "%v is failed"
		starting = "%v is starting"
	}
	type = "slack"
	when = [
		"pipeline.starting",
		"pipeline.complete",
		"pipeline.failed"
	]
}`, i, address, i, i, i)
	}

	return `
resource "spinnaker_pipeline" "test" {
	application = "app"
	name        = "not_pipe"
}` + notifications
}

func testAccCheckPipelineNotifications(resourceName string, expected []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Pipeline not found: %s", resourceName)
		}

		c := testAccProvider.Meta().(*client.Client)
		pipeline, err := c.GetPipelineByID(rs.Primary.Attributes["id"])
		if err != nil {
			return err
		}

		if len(expected) != len(pipeline.Notifications) {
			return fmt.Errorf("Notifications count of %v is expected to be %v",
				len(pipeline.Notifications), len(expected))
		}

		for _, notificationResourceName := range expected {
			expectedResource, ok := s.RootModule().Resources[notificationResourceName]
			if !ok {
				return fmt.Errorf("Notification not found: %s", resourceName)
			}

			err = ensureNotification(pipeline.Notifications, expectedResource)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func ensureNotification(notifications []client.Notification, expected *terraform.ResourceState) error {
	expectedID := expected.Primary.Attributes["id"]
	for _, notification := range notifications {
		if notification.ID == expectedID {
			return nil
		}
	}
	return fmt.Errorf("Notification not found %s", expectedID)
}

func testAccCheckPipelineNotificationDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "spinnaker_pipeline_notification" {
			_, err := c.GetPipelineByID(rs.Primary.Attributes["pipeline"])
			if err == nil {
				return fmt.Errorf("Pipeline notification still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
