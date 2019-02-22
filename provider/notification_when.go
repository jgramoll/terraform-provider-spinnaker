package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type when struct {
	Complete string `mapstructure:"complete"`
	Starting string `mapstructure:"starting"`
	Failed   string `mapstructure:"failed"`
}

func toClientWhen(level client.NotificationLevel, w *when) *[]string {
	clientWhen := []string{}
	// TODO
	if level == client.NotificationLevelPipeline {
		if w.Complete == "1" {
			clientWhen = append(clientWhen, client.PipelineCompleteKey)
		}
		if w.Failed == "1" {
			clientWhen = append(clientWhen, client.PipelineFailedKey)
		}
		if w.Starting == "1" {
			clientWhen = append(clientWhen, client.PipelineStartingKey)
		}
	} else if level == client.NotificationLevelStage {
		if w.Complete == "1" {
			clientWhen = append(clientWhen, client.StageCompleteKey)
		}
		if w.Failed == "1" {
			clientWhen = append(clientWhen, client.StageFailedKey)
		}
		if w.Starting == "1" {
			clientWhen = append(clientWhen, client.StageStartingKey)
		}
	}
	return &clientWhen
}

func (w *when) fromClientWhen(cn *client.Notification) *when {
	newWhen := when{}
	for _, cw := range cn.When {
		switch cw {
		case client.StageCompleteKey:
			newWhen.Complete = "1"
		case client.PipelineCompleteKey:
			newWhen.Complete = "1"
		case client.StageFailedKey:
			newWhen.Failed = "1"
		case client.PipelineFailedKey:
			newWhen.Failed = "1"
		case client.StageStartingKey:
			newWhen.Starting = "1"
		case client.PipelineStartingKey:
			newWhen.Starting = "1"
		}
	}
	return &newWhen
}
