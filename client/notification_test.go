package client

import (
	"testing"
)

func newNotificationTestPipeline() (*Pipeline, *Notification) {
	expected := Notification{
		ID: "notId",
	}
	return &Pipeline{
		Notifications: &[]*Notification{&expected},
	}, &expected
}

func TestGetNotification(t *testing.T) {
	pipeline, expected := newNotificationTestPipeline()
	n, err := pipeline.GetNotification(expected.ID)
	if err != nil {
		t.Fatal(err)
	}
	if n.ID != expected.ID {
		t.Fatal("Not the expected notification")
	}
}

func TestUpdateNotification(t *testing.T) {
	pipeline, expected := newNotificationTestPipeline()

	updateNotification := Notification(*expected)
	updateNotification.Address = "new addr"
	err := pipeline.UpdateNotification(&updateNotification)
	if err != nil {
		t.Fatal(err)
	}
	if (*pipeline.Notifications)[0].Address == expected.Address {
		t.Fatal("Pipeline Notification was not updated")
	}
}

func TestDeleteNotification(t *testing.T) {
	pipeline, expected := newNotificationTestPipeline()
	err := pipeline.DeleteNotification(expected.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(*pipeline.Notifications) != 0 {
		t.Fatal("Pipeline Notification was not deleted")
	}
}
