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
	when = {
		complete = true
		starting = false
		failed = true
	}
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

		pipelineService := testAccProvider.Meta().(*Services).PipelineService
		pipeline, err := pipelineService.GetPipelineByID(rs.Primary.Attributes["id"])
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
			err := ensureMessage(&notification, expected)
			if err != nil {
				return err
			}
			return ensureWhen(&notification, expected)
		}
	}
	return fmt.Errorf("Notification not found %s", expectedID)
}

func ensureMessage(notification *client.Notification, expected *terraform.ResourceState) error {
	if notification.Message.Complete.Text != expected.Primary.Attributes["message.complete"] {
		return fmt.Errorf("Expected complete mesage \"%s\", not \"%s\"",
			expected.Primary.Attributes["message.complete"], notification.Message.Complete.Text)
	}
	if notification.Message.Starting.Text != expected.Primary.Attributes["message.starting"] {
		return fmt.Errorf("Expected starting mesage \"%s\", not \"%s\"",
			expected.Primary.Attributes["message.starting"], notification.Message.Complete.Text)
	}
	if notification.Message.Failed.Text != expected.Primary.Attributes["message.failed"] {
		return fmt.Errorf("Expected failed mesage \"%s\", not \"%s\"",
			expected.Primary.Attributes["message.failed"], notification.Message.Complete.Text)
	}
	return nil
}

func ensureWhen(notification *client.Notification, expected *terraform.ResourceState) error {
	modes := []string{
		"complete",
		"failed",
		"starting",
	}

	for _, mode := range modes {
		expectedWhen := expected.Primary.Attributes[fmt.Sprintf("when.%s", mode)]
		expectedPipeWhen := fmt.Sprintf("pipeline.%s", mode)
		err := whenContainsState(notification.When, expectedPipeWhen)

		if expectedWhen == "1" {
			if err != nil {
				return err
			}
		} else {
			if err == nil {
				return fmt.Errorf("When contained %s, when it should not have", mode)
			}
		}
	}
	return nil
}

func whenContainsState(when []string, expected string) error {
	for _, w := range when {
		if w == expected {
			return nil
		}
	}
	return fmt.Errorf("When not found %s", expected)
}

func testAccCheckPipelineNotificationDestroy(s *terraform.State) error {
	pipelineService := testAccProvider.Meta().(*Services).PipelineService
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "spinnaker_pipeline_notification" {
			_, err := pipelineService.GetPipelineByID(rs.Primary.Attributes["pipeline"])
			if err == nil {
				return fmt.Errorf("Pipeline notification still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
