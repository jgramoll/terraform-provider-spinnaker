package client

import (
	"testing"
)

func TestNewInvalidMessage(t *testing.T) {
	message, err := NewMessage("badLevel")
	if err == nil {
		t.Fatal("Should return error if level is invalid")
	}
	if message != nil {
		t.Fatal("Should not return message")
	}
}

func TestParseMessage(t *testing.T) {
	expectedCompleteText := "is complete"
	expectedFailedText := "is failed"
	messageMap := map[string]interface{}{
		"stage.complete": map[string]string{"text": expectedCompleteText},
		"stage.failed":   map[string]string{"text": expectedFailedText},
	}
	message, err := parseMessage(NotificationLevelStage, messageMap)
	if err != nil {
		t.Fatal(err)
	}
	stageMessage, ok := message.(*StageMessage)
	if !ok {
		t.Fatal("Should return Stage Message")
	}
	if stageMessage.CompleteText() != expectedCompleteText {
		t.Fatalf("Expected complete text %v not %v", expectedCompleteText, stageMessage.CompleteText())
	}
	if stageMessage.FailedText() != expectedFailedText {
		t.Fatalf("Expected failed text %v not %v", expectedFailedText, stageMessage.FailedText())
	}
}
