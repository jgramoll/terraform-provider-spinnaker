package provider

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

// Docker trigger for Pipeline
type dockerTrigger struct {
	Enabled   bool   `mapstructure:"enabled"`
	RunAsUser string `mapstructure:"run_as_user"`

	Account      string `mapstructure:"account"`
	Organization string `mapstructure:"organization"`
	Registry     string `mapstructure:"registry"`
	Repository   string `mapstructure:"repository"`
	Tag          string `mapstructure:"tag"`
}

func newDockerTrigger() *dockerTrigger {
	return &dockerTrigger{}
}

func (t *dockerTrigger) toClientTrigger(id string) (client.Trigger, error) {
	clientTrigger := client.NewDockerTrigger()
	clientTrigger.ID = id
	clientTrigger.Enabled = t.Enabled
	clientTrigger.RunAsUser = t.RunAsUser

	clientTrigger.Account = t.Account
	clientTrigger.Organization = t.Organization
	clientTrigger.Registry = t.Registry
	clientTrigger.Repository = t.Repository
	clientTrigger.Tag = t.Tag
	return clientTrigger, nil
}

func (*dockerTrigger) fromClientTrigger(clientTriggerInterface client.Trigger) (trigger, error) {
	clientTrigger, ok := clientTriggerInterface.(*client.DockerTrigger)
	if !ok {
		return nil, errors.New("Expected docker trigger")
	}
	t := newDockerTrigger()
	t.RunAsUser = clientTrigger.RunAsUser
	t.Enabled = clientTrigger.Enabled

	t.Account = clientTrigger.Account
	t.Organization = clientTrigger.Organization
	t.Registry = clientTrigger.Registry
	t.Repository = clientTrigger.Repository
	t.Tag = clientTrigger.Tag
	return t, nil
}

func (t *dockerTrigger) setResourceData(d *schema.ResourceData) error {
	var err error
	err = d.Set("enabled", t.Enabled)
	if err != nil {
		return err
	}
	err = d.Set("run_as_user", t.RunAsUser)
	if err != nil {
		return err
	}
	err = d.Set("account", t.Account)
	if err != nil {
		return err
	}
	err = d.Set("organization", t.Organization)
	if err != nil {
		return err
	}
	err = d.Set("registry", t.Registry)
	if err != nil {
		return err
	}
	err = d.Set("repository", t.Repository)
	if err != nil {
		return err
	}
	err = d.Set("tag", t.Tag)
	if err != nil {
		return err
	}
	return nil
}
