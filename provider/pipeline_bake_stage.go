package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type bakeStage struct {
	baseStage `mapstructure:",squash"`

	AmiName            string            `mapstructure:"ami_name"`
	AmiSuffix          string            `mapstructure:"ami_suffix"`
	BaseAMI            string            `mapstructure:"base_ami"`
	BaseLabel          string            `mapstructure:"base_label"`
	BaseName           string            `mapstructure:"base_name"`
	BaseOS             string            `mapstructure:"base_os"`
	CloudProviderType  string            `mapstructure:"cloud_provider_type"`
	ExtendedAttributes map[string]string `mapstructure:"extended_attributes"`
	Package            string            `mapstructure:"package"`
	Rebake             bool              `mapstructure:"rebake"`
	Region             string            `mapstructure:"region"`
	Regions            []string          `mapstructure:"regions"`
	StoreType          string            `mapstructure:"store_type"`
	TemplateFileName   string            `mapstructure:"template_file_name"`
	User               string            `mapstructure:"user"`
	VarFileName        string            `mapstructure:"var_file_name"`
	VMType             string            `mapstructure:"vm_type"`
}

func newBakeStage() *bakeStage {
	return &bakeStage{
		baseStage: *newBaseStage(),
	}
}

func (s *bakeStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewBakeStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.AmiName = s.AmiName
	cs.AmiSuffix = s.AmiSuffix
	cs.BaseAMI = s.BaseAMI
	cs.BaseLabel = s.BaseLabel
	cs.BaseName = s.BaseName
	cs.BaseOS = s.BaseOS
	cs.CloudProviderType = s.CloudProviderType
	cs.ExtendedAttributes = s.ExtendedAttributes
	cs.Package = s.Package
	cs.Rebake = s.Rebake
	if s.Region == "" && len(s.Regions) == 1 {
		cs.Region = s.Regions[0]
	} else {
		cs.Region = s.Region
	}
	cs.Regions = s.Regions
	cs.StoreType = s.StoreType
	cs.TemplateFileName = s.TemplateFileName
	if s.User == "" {
		cs.User = config.Auth.UserEmail
	} else {
		cs.User = s.User
	}
	cs.VarFileName = s.VarFileName
	cs.VMType = s.VMType

	return cs, nil
}

func (*bakeStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.BakeStage)
	newStage := newBakeStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.AmiName = clientStage.AmiName
	newStage.AmiSuffix = clientStage.AmiSuffix
	newStage.BaseAMI = clientStage.BaseAMI
	newStage.BaseLabel = clientStage.BaseLabel
	newStage.BaseName = clientStage.BaseName
	newStage.BaseOS = clientStage.BaseOS
	newStage.CloudProviderType = clientStage.CloudProviderType
	newStage.ExtendedAttributes = clientStage.ExtendedAttributes
	newStage.Package = clientStage.Package
	newStage.Rebake = clientStage.Rebake
	newStage.Region = clientStage.Region
	newStage.Regions = clientStage.Regions
	newStage.StoreType = clientStage.StoreType
	newStage.TemplateFileName = clientStage.TemplateFileName
	newStage.User = clientStage.User
	newStage.VarFileName = clientStage.VarFileName
	newStage.VMType = clientStage.VMType

	return newStage, nil
}

func (s *bakeStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("ami_name", s.AmiName)
	if err != nil {
		return err
	}
	err = d.Set("ami_suffix", s.AmiSuffix)
	if err != nil {
		return err
	}
	err = d.Set("base_ami", s.BaseAMI)
	if err != nil {
		return err
	}
	err = d.Set("base_label", s.BaseLabel)
	if err != nil {
		return err
	}
	err = d.Set("base_name", s.BaseName)
	if err != nil {
		return err
	}
	err = d.Set("base_os", s.BaseOS)
	if err != nil {
		return err
	}
	err = d.Set("cloud_provider_type", s.CloudProviderType)
	if err != nil {
		return err
	}
	err = d.Set("extended_attributes", s.ExtendedAttributes)
	if err != nil {
		return err
	}
	err = d.Set("package", s.Package)
	if err != nil {
		return err
	}
	err = d.Set("rebake", s.Rebake)
	if err != nil {
		return err
	}
	err = d.Set("regions", s.Regions)
	if err != nil {
		return err
	}
	err = d.Set("store_type", s.StoreType)
	if err != nil {
		return err
	}
	err = d.Set("template_file_name", s.TemplateFileName)
	if err != nil {
		return err
	}
	err = d.Set("var_file_name", s.VarFileName)
	if err != nil {
		return err
	}
	return d.Set("vm_type", s.VMType)
}
