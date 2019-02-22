package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type stage interface {
	fromClientStage(client.Stage) stage
	toClientStage(*client.Config) (client.Stage, error)
	SetResourceData(*schema.ResourceData) error
	SetRefID(string)
	GetRefID() string
}

// TODO why does this not like mapstructure
// type baseStage struct {
// 	Name  string           `mapstructure:"name"`
// 	RefID string           `mapstructure:"ref_id"`
// 	Type  client.StageType `mapstructure:"type"`
// }
