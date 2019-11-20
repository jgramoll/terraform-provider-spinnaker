package client

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestNewPipeline(t *testing.T) {
	pipeline := NewPipeline()
	if pipeline.Disabled {
		t.Fatal("default pipeline should not be disabled")
	}
}

func TestParseEmptyPipeline(t *testing.T) {
	pipeline, err := parsePipeline(nil)
	if err != nil {
		t.Fatal(err)
	}
	if pipeline.Disabled {
		t.Fatal("default pipeline should not be disabled")
	}
}

func TestParsePipeline(t *testing.T) {
	expectedName := "test-expected"

	expectedStage := NewBakeStage()
	expectedStage.Name = "test-stage"
	expectedStage.AmiName = "ami-test"

	expectedNotification := Notification{
		SerializableNotification: SerializableNotification{
			Address: "test-address",
			Level:   NotificationLevelPipeline,
			Type:    "slack",
			When: []string{
				PipelineStartingKey,
				PipelineCompleteKey,
			},
		},
		Message: &PipelineMessage{
			Complete: &MessageText{Text: "pipe is complete"},
			Failed:   &MessageText{Text: "pipe is failed"},
		},
	}
	expectedTriggers := []Trigger{
		NewJenkinsTrigger(),
		NewPipelineTrigger(),
	}

	pipeline, err := parsePipeline(map[string]interface{}{
		"name": expectedName,
		"stages": []interface{}{
			map[string]interface{}{
				"type":    BakeStageType.String(),
				"name":    expectedStage.Name,
				"amiName": expectedStage.AmiName,
			},
		},
		"notifications": []interface{}{
			map[string]interface{}{
				"address": expectedNotification.Address,
				"level":   expectedNotification.Level,
				"message": map[string]interface{}{
					PipelineCompleteKey: map[string]string{
						"text": expectedNotification.Message.CompleteText(),
					},
					PipelineFailedKey: map[string]string{
						"text": expectedNotification.Message.FailedText(),
					},
				},
				"type": expectedNotification.Type,
				"when": expectedNotification.When,
			},
		},
		"triggers": []interface{}{
			map[string]interface{}{
				"type": JenkinsTriggerType.String(),
			},
			map[string]interface{}{
				"type": PipelineTriggerType.String(),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	expected := NewPipeline()
	expected.Name = expectedName
	expected.Stages = &[]Stage{expectedStage}
	expected.Notifications = &[]*Notification{&expectedNotification}
	expected.Triggers = expectedTriggers
	err = pipeline.equals(expected)
	if err != nil {
		t.Fatal(err)
	}
}

func (pipeline *Pipeline) equals(expected *Pipeline) error {
	if pipeline.Notifications != nil {
		if expected.Notifications == nil {
			return errors.New("Expected Notifications, but was nil")
		}
		for i, n := range *pipeline.Notifications {
			expectedNotifications := *expected.Notifications
			if !reflect.DeepEqual(n.Message, expectedNotifications[i].Message) {
				return fmt.Errorf("Pipeline Notification Message %v does not match %v", n.Message, expectedNotifications[i].Message)
			}
			if !reflect.DeepEqual(n, expectedNotifications[i]) {
				return fmt.Errorf("Pipeline Notification %v does not match %v", n, expectedNotifications[i])
			}
		}
	}

	if !reflect.DeepEqual(pipeline, expected) {
		return fmt.Errorf("Pipeline %v does not match %v", pipeline, expected)
	}
	return nil
}
