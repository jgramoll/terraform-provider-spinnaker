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

func TestAccPipelineNotificationBasic(t *testing.T) {
	var pipelineRef client.Pipeline
	var notifications []*client.Notification
	pipeName := fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	address := "bridge-career-deploys"
	addressChanged := address + "-new"
	pipelineResourceName := "spinnaker_pipeline.test"
	notification1 := "spinnaker_pipeline_notification.n1"
	notification2 := "spinnaker_pipeline_notification.n2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineNotificationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPipelineNotificationConfigBasic(pipeName, address, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(notification1, "address", address),
					resource.TestCheckResourceAttr(notification1, "message.0.complete", "1 is done"),
					resource.TestCheckResourceAttr(notification2, "address", address),
					resource.TestCheckResourceAttr(notification2, "message.0.complete", "2 is done"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineNotifications(pipelineResourceName, []string{
						notification1,
						notification2,
					}, &notifications),
				),
			},
			{
				ResourceName:  notification1,
				ImportStateId: "invalid",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import key, must be pipelineID_notificationID`),
			},
			{
				ResourceName: notification1,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(notifications) == 0 {
						return "", fmt.Errorf("no notifications to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, notifications[0].ID), nil
				},
				ImportStateVerify: true,
			},
			{
				ResourceName: notification2,
				ImportState:  true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					if len(notifications) < 2 {
						return "", fmt.Errorf("no notifications to import")
					}
					return fmt.Sprintf("%s_%s", pipelineRef.ID, notifications[1].ID), nil
				},
				ImportStateVerify: true,
			},
			{
				Config: testAccPipelineNotificationConfigBasic(pipeName, addressChanged, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(notification1, "address", addressChanged),
					resource.TestCheckResourceAttr(notification1, "message.0.complete", "1 is done"),
					resource.TestCheckResourceAttr(notification2, "address", addressChanged),
					resource.TestCheckResourceAttr(notification2, "message.0.complete", "2 is done"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineNotifications(pipelineResourceName, []string{
						notification1,
						notification2,
					}, &notifications),
				),
			},
			{
				Config: testAccPipelineNotificationConfigBasic(pipeName, address, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(notification1, "address", address),
					resource.TestCheckResourceAttr(notification1, "message.0.complete", "1 is done"),
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineNotifications(pipelineResourceName, []string{
						notification1,
					}, &notifications),
				),
			},
			{
				Config: testAccPipelineNotificationConfigBasic(pipeName, address, 0),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPipelineExists(pipelineResourceName, &pipelineRef),
					testAccCheckPipelineNotifications(pipelineResourceName, []string{}, &notifications),
				),
			},
		},
	})
}

func testAccPipelineNotificationConfigBasic(pipeName string, address string, count int) string {
	notifications := ""
	for i := 1; i <= count; i++ {
		notifications += fmt.Sprintf(`
resource "spinnaker_pipeline_notification" "n%v" {
	pipeline = "${spinnaker_pipeline.test.id}"
	address = "%v"
	message {
		complete = "%v is done"
		failed = "%v is failed"
		starting = "%v is starting"
	}
	type = "slack"
	when {
		complete = true
		starting = false
		failed = true
	}
}`, i, address, i, i, i)
	}

	return testAccPipelineConfigBasic("app", pipeName) + notifications
}

func testAccCheckPipelineNotifications(resourceName string, expected []string, notifications *[]*client.Notification) resource.TestCheckFunc {
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

		if len(expected) != len(*pipeline.Notifications) {
			return fmt.Errorf("Notifications count of %v is expected to be %v",
				len(*pipeline.Notifications), len(expected))
		}

		for _, notificationResourceName := range expected {
			expectedResource, ok := s.RootModule().Resources[notificationResourceName]
			if !ok {
				return fmt.Errorf("Notification not found: %s", resourceName)
			}

			notification, err := ensureNotification(pipeline.Notifications, expectedResource)
			if err != nil {
				return err
			}
			*notifications = append(*notifications, notification)
		}

		return nil
	}
}

func ensureNotification(notifications *[]*client.Notification, expected *terraform.ResourceState) (*client.Notification, error) {
	expectedID := expected.Primary.Attributes["id"]
	for _, notification := range *notifications {
		if notification.ID == expectedID {
			err := ensureMessage(notification, expected)
			if err != nil {
				return nil, err
			}
			err = ensureWhen(notification, expected)
			if err != nil {
				return nil, err
			}
			return notification, nil
		}
	}
	return nil, fmt.Errorf("Notification not found %s", expectedID)
}

func ensureMessage(notification *client.Notification, expected *terraform.ResourceState) error {
	if notification.Message.CompleteText() != expected.Primary.Attributes["message.0.complete"] {
		return fmt.Errorf("Expected complete mesage \"%s\", not \"%s\"",
			expected.Primary.Attributes["message.0.complete"], notification.Message.CompleteText())
	}
	if notification.Message.StartingText() != expected.Primary.Attributes["message.0.starting"] {
		return fmt.Errorf("Expected starting mesage \"%s\", not \"%s\"",
			expected.Primary.Attributes["message.0.starting"], notification.Message.StartingText())
	}
	if notification.Message.FailedText() != expected.Primary.Attributes["message.0.failed"] {
		return fmt.Errorf("Expected failed mesage \"%s\", not \"%s\"",
			expected.Primary.Attributes["message.0.failed"], notification.Message.FailedText())
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
		expectedWhen := expected.Primary.Attributes[fmt.Sprintf("when.0.%s", mode)]
		expectedPipeWhen := fmt.Sprintf("pipeline.%s", mode)
		err := whenContainsState(notification.When, expectedPipeWhen)

		if expectedWhen == "true" {
			if err != nil {
				return err
			}
		} else {
			if err == nil {
				return fmt.Errorf("When contained %s, when it should not have. %v %v", mode, expectedWhen, notification.When)
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
			_, err := pipelineService.GetPipelineByID(rs.Primary.Attributes[PipelineKey])
			if err == nil {
				return fmt.Errorf("Pipeline notification still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
