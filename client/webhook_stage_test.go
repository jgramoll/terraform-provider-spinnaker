package client

import (
	"testing"
)

var webhookStage WebhookStage

func init() {
	webhookStage = *NewWebhookStage()
}

func TestNewWebhookStage(t *testing.T) {
	if webhookStage.Type != WebhookStageType {
		t.Fatalf("Deploy stage type should be %s, not \"%s\"", WebhookStageType, webhookStage.Type)
	}
}

func TestWebhookStageGetName(t *testing.T) {
	name := "New Deploy"
	webhookStage.Name = name
	if webhookStage.GetName() != name {
		t.Fatalf("Deploy stage GetName() should be %s, not \"%s\"", name, webhookStage.GetName())
	}
}

func TestWebhookStageGetType(t *testing.T) {
	if webhookStage.GetType() != WebhookStageType {
		t.Fatalf("Deploy stage GetType() should be %s, not \"%s\"", WebhookStageType, webhookStage.GetType())
	}
	if webhookStage.Type != WebhookStageType {
		t.Fatalf("Deploy stage Type should be %s, not \"%s\"", WebhookStageType, webhookStage.Type)
	}
}
