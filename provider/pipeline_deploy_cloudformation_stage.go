package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type deployCloudformationStage struct {
	baseStage `mapstructure:",squash"`

	ActionOnReplacement string            `mapstructure:"action_on_replacement"`
	Capabilities        []string          `mapstructure:"capabilities"`
	ChangeSetName       string            `mapstructure:"change_set_name"`
	Credentials         string            `mapstructure:"credentials"`
	ExecuteChangeSet    bool              `mapstructure:"execute_change_set"`
	IsChangeSet         bool              `mapstructure:"is_change_set"`
	Parameters          map[string]string `mapstructure:"parameters"`
	Regions             []string          `mapstructure:"regions"`
	RoleARN             string            `mapstructure:"role_arn"`
	Source              string            `mapstructure:"source"`
	StackArtifact       *[]*stackArtifact `mapstructure:"stack_artifact"`
	StackName           string            `mapstructure:"stack_name"`
	Tags                map[string]string `mapstructure:"tags"`
	TemplateBody        []string          `mapstructure:"template_body"`
}

func newDeployCloudformationStage() *deployCloudformationStage {
	return &deployCloudformationStage{
		baseStage: *newBaseStage(),
		Source:    "text",
	}
}

func (s *deployCloudformationStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewDeployCloudformationStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.ActionOnReplacement = s.ActionOnReplacement
	cs.Capabilities = s.Capabilities
	cs.ChangeSetName = s.ChangeSetName
	cs.Credentials = s.Credentials
	cs.ExecuteChangeSet = s.ExecuteChangeSet
	cs.IsChangeSet = s.IsChangeSet
	cs.Parameters = s.Parameters
	cs.Regions = s.Regions
	cs.RoleARN = s.RoleARN
	cs.StackArtifact = toClientStackArtifact(s.StackArtifact)
	cs.StackName = s.StackName
	cs.Tags = s.Tags
	cs.TemplateBody = s.TemplateBody

	return cs, nil
}

func (*deployCloudformationStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.DeployCloudformationStage)
	newStage := newDeployCloudformationStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.ActionOnReplacement = clientStage.ActionOnReplacement
	newStage.Capabilities = clientStage.Capabilities
	newStage.ChangeSetName = clientStage.ChangeSetName
	newStage.Credentials = clientStage.Credentials
	newStage.ExecuteChangeSet = clientStage.ExecuteChangeSet
	newStage.IsChangeSet = clientStage.IsChangeSet
	newStage.Parameters = clientStage.Parameters
	newStage.Regions = clientStage.Regions
	newStage.RoleARN = clientStage.RoleARN
	newStage.StackArtifact = fromClientStackArtifact(clientStage.StackArtifact)
	newStage.StackName = clientStage.StackName
	newStage.Tags = clientStage.Tags
	newStage.TemplateBody = clientStage.TemplateBody

	return newStage, nil
}

func (s *deployCloudformationStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("action_on_replacement", s.ActionOnReplacement)
	if err != nil {
		return err
	}
	err = d.Set("capabilities", s.Capabilities)
	if err != nil {
		return err
	}
	err = d.Set("change_set_name", s.ChangeSetName)
	if err != nil {
		return err
	}
	err = d.Set("credentials", s.Credentials)
	if err != nil {
		return err
	}
	err = d.Set("execute_change_set", s.ExecuteChangeSet)
	if err != nil {
		return err
	}
	err = d.Set("is_change_set", s.IsChangeSet)
	if err != nil {
		return err
	}
	err = d.Set("parameters", s.Parameters)
	if err != nil {
		return err
	}
	err = d.Set("regions", s.Regions)
	if err != nil {
		return err
	}
	err = d.Set("role_arn", s.RoleARN)
	if err != nil {
		return err
	}
	err = d.Set("source", s.Source)
	if err != nil {
		return err
	}
	err = d.Set("stack_artifact", s.StackArtifact)
	if err != nil {
		return err
	}
	err = d.Set("stack_name", s.StackName)
	if err != nil {
		return err
	}
	err = d.Set("tags", s.Tags)
	if err != nil {
		return err
	}
	err = d.Set("template_body", s.TemplateBody)
	if err != nil {
		return err
	}

	return nil
}
