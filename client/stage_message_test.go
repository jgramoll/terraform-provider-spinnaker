package client

import (
	"testing"
)

func TestNewStageMessage(t *testing.T) {
	message, err := NewMessage(NotificationLevelStage)
	if err != nil {
		t.Fatal(err)
	}
	_, ok := message.(*StageMessage)
	if !ok {
		t.Fatal("Should return StageMessage")
	}
}
