package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type bakeStage struct {
	baseStage

	AmiName            string            `mapstructure:"ami_name"`
	AmiSuffix          string            `mapstructure:"ami_suffix"`
	BaseAMI            string            `mapstructure:"base_ami"`
	BaseLabel          string            `mapstructure:"base_label"`
	BaseName           string            `mapstructure:"base_name"`
	BaseOS             string            `mapstructure:"base_os"`
	CloudProviderType  string            `mapstructure:"cloud_provider_type"`
	ExtendedAttributes map[string]string `mapstructure:"extended_attributes"`
	Regions            []string          `mapstructure:"regions"`
	RequisiteStages    []string          `mapstructure:"requisite_stages"`
	StoreType          string            `mapstructure:"store_type"`
	TemplateFileName   string            `mapstructure:"template_file_name"`
	VarFileName        string            `mapstructure:"var_file_name"`
	VMType             string            `mapstructure:"vm_type"`
}

func newBakeStage() stage {
	return bakeStage{baseStage: baseStage{Type: client.BakeStageType}}
}

func (s bakeStage) toClientStage() client.Stage {
	return client.BakeStage{
		// BaseStage: client.BaseStage{
		// TODO
		Name: s.Name,
		// },
	}
}

func (s bakeStage) SetResourceData() {
	// TODO
}
