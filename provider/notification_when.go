package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type when struct {
	Complete bool `mapstructure:"complete"`
	Starting bool `mapstructure:"starting"`
	Failed   bool `mapstructure:"failed"`
}

func newWhen() *when {
	return &when{}
}

func toClientWhen(level client.NotificationLevel, w *when) *[]string {
	clientWhen := []string{}
	if w == nil {
		return &clientWhen
	}
	if level == client.NotificationLevelPipeline {
		if w.Complete {
			clientWhen = append(clientWhen, client.PipelineCompleteKey)
		}
		if w.Failed {
			clientWhen = append(clientWhen, client.PipelineFailedKey)
		}
		if w.Starting {
			clientWhen = append(clientWhen, client.PipelineStartingKey)
		}
	} else if level == client.NotificationLevelStage {
		if w.Complete {
			clientWhen = append(clientWhen, client.StageCompleteKey)
		}
		if w.Failed {
			clientWhen = append(clientWhen, client.StageFailedKey)
		}
		if w.Starting {
			clientWhen = append(clientWhen, client.StageStartingKey)
		}
	}
	return &clientWhen
}

func fromClientWhen(cn *client.Notification) *when {
	w := newWhen()
	for _, cw := range cn.When {
		switch cw {
		case client.StageCompleteKey:
			w.Complete = true
		case client.PipelineCompleteKey:
			w.Complete = true
		case client.StageFailedKey:
			w.Failed = true
		case client.PipelineFailedKey:
			w.Failed = true
		case client.StageStartingKey:
			w.Starting = true
		case client.PipelineStartingKey:
			w.Starting = true
		}
	}
	return w
}
