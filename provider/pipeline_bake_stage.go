package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type bakeStage struct {
	// baseStage
	Name                              string                   `mapstructure:"name"`
	RefID                             string                   `mapstructure:"ref_id"`
	Type                              client.StageType         `mapstructure:"type"`
	RequisiteStageRefIds              []string                 `mapstructure:"requisite_stage_ref_ids"`
	Notifications                     *[]*notification         `mapstructure:"notification"`
	StageEnabled                      *[]*stageEnabled         `mapstructure:"stage_enabled"`
	CompleteOtherBranchesThenFail     bool                     `mapstructure:"complete_other_branches_then_fail"`
	ContinuePipeline                  bool                     `mapstructure:"continue_pipeline"`
	FailOnFailedExpressions           bool                     `mapstructure:"fail_on_failed_expressions"`
	FailPipeline                      bool                     `mapstructure:"fail_pipeline"`
	OverrideTimeout                   bool                     `mapstructure:"override_timeout"`
	RestrictExecutionDuringTimeWindow bool                     `mapstructure:"restrict_execution_during_time_window"`
	RestrictedExecutionWindow         *[]*stageExecutionWindow `mapstructure:"restricted_execution_window"`
	// End baseStage

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
		Type:         client.BakeStageType,
		FailPipeline: true,
	}
}

func (s *bakeStage) toClientStage(config *client.Config) (client.Stage, error) {
	// baseStage
	notifications, err := toClientNotifications(s.Notifications)
	if err != nil {
		return nil, err
	}

	cs := client.NewBakeStage()
	cs.Name = s.Name
	cs.RefID = s.RefID
	cs.RequisiteStageRefIds = s.RequisiteStageRefIds
	cs.Notifications = notifications
	cs.SendNotifications = notifications != nil && len(*notifications) > 0
	cs.StageEnabled = toClientStageEnabled(s.StageEnabled)
	cs.CompleteOtherBranchesThenFail = s.CompleteOtherBranchesThenFail
	cs.ContinuePipeline = s.ContinuePipeline
	cs.FailOnFailedExpressions = s.FailOnFailedExpressions
	cs.FailPipeline = s.FailPipeline
	cs.OverrideTimeout = s.OverrideTimeout
	cs.RestrictExecutionDuringTimeWindow = s.RestrictExecutionDuringTimeWindow
	cs.RestrictedExecutionWindow = toClientExecutionWindow(s.RestrictedExecutionWindow)
	// End baseStage

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

func (s *bakeStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.BakeStage)
	newStage := newBakeStage()

	// baseStage
	newStage.Name = clientStage.Name
	newStage.RefID = clientStage.RefID
	newStage.RequisiteStageRefIds = clientStage.RequisiteStageRefIds
	newStage.Notifications = fromClientNotifications(clientStage.Notifications)
	newStage.StageEnabled = fromClientStageEnabled(clientStage.StageEnabled)
	newStage.CompleteOtherBranchesThenFail = clientStage.CompleteOtherBranchesThenFail
	newStage.ContinuePipeline = clientStage.ContinuePipeline
	newStage.FailOnFailedExpressions = clientStage.FailOnFailedExpressions
	newStage.FailPipeline = clientStage.FailPipeline
	newStage.OverrideTimeout = clientStage.OverrideTimeout
	newStage.RestrictExecutionDuringTimeWindow = clientStage.RestrictExecutionDuringTimeWindow
	newStage.RestrictedExecutionWindow = fromClientExecutionWindow(clientStage.RestrictedExecutionWindow)
	// end baseStage

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

	return newStage
}

func (s *bakeStage) SetResourceData(d *schema.ResourceData) error {
	// baseStage
	err := d.Set("name", s.Name)
	if err != nil {
		return err
	}
	err = d.Set("requisite_stage_ref_ids", s.RequisiteStageRefIds)
	if err != nil {
		return err
	}
	err = d.Set("notification", s.Notifications)
	if err != nil {
		return err
	}
	err = d.Set("stage_enabled", s.StageEnabled)
	if err != nil {
		return err
	}
	err = d.Set("complete_other_branches_then_fail", s.CompleteOtherBranchesThenFail)
	if err != nil {
		return err
	}
	err = d.Set("continue_pipeline", s.ContinuePipeline)
	if err != nil {
		return err
	}
	err = d.Set("fail_on_failed_expressions", s.FailOnFailedExpressions)
	if err != nil {
		return err
	}
	err = d.Set("fail_pipeline", s.FailPipeline)
	if err != nil {
		return err
	}
	err = d.Set("override_timeout", s.OverrideTimeout)
	if err != nil {
		return err
	}
	err = d.Set("restrict_execution_during_time_window", s.RestrictExecutionDuringTimeWindow)
	if err != nil {
		return err
	}
	err = d.Set("restricted_execution_window", s.RestrictedExecutionWindow)
	if err != nil {
		return err
	}
	// End baseStage

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

func (s *bakeStage) SetRefID(id string) {
	s.RefID = id
}

func (s *bakeStage) GetRefID() string {
	return s.RefID
}
