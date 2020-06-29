package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type scriptStage struct {
	baseStage `mapstructure:",squash"`

	Account           string   `mapstructure:"account"`
	Cluster           string   `mapstructure:"cluster"`
	Clusters          []string `mapstructure:"clusters"`
	Command           string   `mapstructure:"command"`
	Cmc               string   `mapstructure:"cmc"`
	Image             string   `mapstructure:"image"`
	PropertyFile      string   `mapstructure:"property_file"`
	Region            string   `mapstructure:"region"`
	Regions           []string `mapstructure:"regions"`
	RepoURL           string   `mapstructure:"repo_url"`
	RepoBranch        string   `mapstructure:"repo_branch"`
	ScriptPath        string   `mapstructure:"script_path"`
	WaitForCompletion bool     `mapstructure:"wait_for_completion"`
}

func newScriptStage() *scriptStage {
	return &scriptStage{
		baseStage: *newBaseStage(),
	}
}

func (s *scriptStage) toClientStage(config *client.Config, refID string) (client.Stage, error) {
	cs := client.NewScriptStage()
	err := s.baseToClientStage(&cs.BaseStage, refID, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	cs.Account = s.Account
	if s.Cluster == "" && len(s.Clusters) == 1 {
		cs.Cluster = s.Clusters[0]
	} else {
		cs.Cluster = s.Cluster
	}
	cs.Clusters = s.Clusters
	cs.Command = s.Command
	cs.Cmc = s.Cmc
	cs.Image = s.Image
	cs.PropertyFile = s.PropertyFile
	if s.Region == "" && len(s.Regions) == 1 {
		cs.Region = s.Regions[0]
	} else {
		cs.Region = s.Region
	}
	cs.Regions = s.Regions
	cs.RepoURL = s.RepoURL
	cs.RepoBranch = s.RepoBranch
	cs.ScriptPath = s.ScriptPath
	cs.WaitForCompletion = s.WaitForCompletion

	return cs, nil
}

func (*scriptStage) fromClientStage(cs client.Stage) (stage, error) {
	clientStage := cs.(*client.ScriptStage)
	newStage := newScriptStage()
	err := newStage.baseFromClientStage(&clientStage.BaseStage, newDefaultNotificationInterface)
	if err != nil {
		return nil, err
	}

	newStage.Account = clientStage.Account
	newStage.Cluster = clientStage.Cluster
	newStage.Clusters = clientStage.Clusters
	newStage.Command = clientStage.Command
	newStage.Cmc = clientStage.Cmc
	newStage.Image = clientStage.Image
	newStage.PropertyFile = clientStage.PropertyFile
	newStage.Region = clientStage.Region
	newStage.Regions = clientStage.Regions
	newStage.RepoURL = clientStage.RepoURL
	newStage.RepoBranch = clientStage.RepoBranch
	newStage.ScriptPath = clientStage.ScriptPath
	newStage.WaitForCompletion = clientStage.WaitForCompletion

	return newStage, nil
}

func (s *scriptStage) SetResourceData(d *schema.ResourceData) error {
	err := s.baseSetResourceData(d)
	if err != nil {
		return err
	}

	err = d.Set("account", s.Account)
	if err != nil {
		return err
	}

	err = d.Set("cluster", s.Cluster)
	if err != nil {
		return err
	}

	err = d.Set("clusters", s.Clusters)
	if err != nil {
		return err
	}

	err = d.Set("command", s.Command)
	if err != nil {
		return err
	}

	err = d.Set("cmc", s.Cmc)
	if err != nil {
		return err
	}

	err = d.Set("image", s.Image)
	if err != nil {
		return err
	}

	err = d.Set("property_file", s.PropertyFile)
	if err != nil {
		return err
	}

	err = d.Set("region", s.Region)
	if err != nil {
		return err
	}

	err = d.Set("regions", s.Regions)
	if err != nil {
		return err
	}

	err = d.Set("repo_url", s.RepoURL)
	if err != nil {
		return err
	}

	err = d.Set("repo_branch", s.RepoBranch)
	if err != nil {
		return err
	}

	err = d.Set("script_path", s.ScriptPath)
	if err != nil {
		return err
	}

	return d.Set("wait_for_completion", s.WaitForCompletion)
}
