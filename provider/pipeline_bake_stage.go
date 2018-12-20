package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type bakeStage struct {
	// TODO does this baseStage work?
	// baseStage
	Name  string           `mapstructure:"name"`
	RefID string           `mapstructure:"ref_id"`
	Type  client.StageType `mapstructure:"type"`

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

func newBakeStage() interface{} {
	return &bakeStage{Type: client.BakeStageType}
}

func (s bakeStage) toClientStage() client.Stage {
	return client.BakeStage(s)
}

// TODO can we just update the ptr?
func (s *bakeStage) fromClientStage(cs client.Stage) stage {
	newStage := bakeStage(*(cs.(*client.BakeStage)))
	// s = &newStage
	return stage(&newStage)
}

func (s *bakeStage) SetResourceData(d *schema.ResourceData) {
	// TODO
	d.Set("name", s.Name)
}

func (s *bakeStage) SetRefID(id string) {
	s.RefID = id
}

func (s *bakeStage) GetRefID() string {
	return s.RefID
}
