package client

import (
	"testing"
)

func TestNewPipelineMessage(t *testing.T) {
	message, err := NewMessage(NotificationLevelPipeline)
	if err != nil {
		t.Fatal(err)
	}
	_, ok := message.(*PipelineMessage)
	if !ok {
		t.Fatal("Should return PipelineMessage")
	}
}
