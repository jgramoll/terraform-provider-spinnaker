package client

import (
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
		Address: "test-address",
		Level:   "pipeline",
		Type:    "slack",
		Message: Message{
			Complete: MessageText{Text: "pipe is complete"},
			Failed:   MessageText{Text: "pipe is failed"},
		},
		When: []string{
			PipelineStartingKey,
			PipelineCompleteKey,
		},
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
						"text": expectedNotification.Message.Complete.Text,
					},
					PipelineFailedKey: map[string]string{
						"text": expectedNotification.Message.Failed.Text,
					},
				},
				"type": expectedNotification.Type,
				"when": expectedNotification.When,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	expected := NewPipeline()
	expected.Name = expectedName
	expected.Stages = []Stage{expectedStage}
	expected.Notifications = []Notification{expectedNotification}
	err = pipeline.equals(expected)
	if err != nil {
		t.Fatal(err)
	}
}

func (pipeline *Pipeline) equals(expected *Pipeline) error {
	if !reflect.DeepEqual(pipeline, expected) {
		return fmt.Errorf("Pipeline %v does not match %v", pipeline, expected)
	}
	return nil
}
