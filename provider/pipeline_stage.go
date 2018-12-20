package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type stage interface {
	fromClientStage(client.Stage) stage
	toClientStage() client.Stage
	SetResourceData(*schema.ResourceData)
	SetRefID(string)
	GetRefID() string
}

type baseStage struct {
	Name  string           `mapstructure:"name"`
	RefID string           `mapstructure:"ref_id"`
	Type  client.StageType `mapstructure:"type"`
}
