package provider

import (
	"strings"

	"github.com/get-bridge/terraform-provider-spinnaker/client"
	"github.com/hashicorp/terraform/helper/schema"
)

// Application deploy application in application
type application struct {
	ID string `mapstructure:"id"`

	Accounts       []string `mapstructure:"accounts"`
	CloudProviders []string `mapstructure:"cloud_providers"`
	// DataSources    *DataSources `mapstructure:"data_sources"`
	// DesiredCount string `json:"desiredCount"`
	Email string `mapstructure:"email"`

	EnableRestartRunningExecutions bool `mapstructure:"enable_restart_running_executions"`

	IamRole      string `mapstructure:"iam_role"`
	InstancePort int    `mapstructure:"instance_port"`
	Name         string `mapstructure:"name"`

	Permissions                    *[]applicationPermissions `mapstructure:"permissions"`
	PlatformHealthOnly             bool                      `mapstructure:"platform_health_only"`
	PlatformHealthOnlyShowOverride bool                      `mapstructure:"platform_health_only_show_override"`
	ProviderSettings               *[]providerSettings       `mapstructure:"provider_settings"`

	RepoProjectKey string   `mapstructure:"repo_project_key"`
	RepoSlug       string   `mapstructure:"repo_slug"`
	RepoType       string   `mapstructure:"repo_type"`
	TaskDefinition string   `mapstructure:"task_definition"`
	TrafficGuards  []string `mapstructure:"traffic_guards"`
}

func (a *application) toClientApplication() *client.Application {
	return &client.Application{
		Accounts:       strings.Join(a.Accounts, ","),
		CloudProviders: strings.Join(a.CloudProviders, ","),
		Email:          a.Email,

		EnableRestartRunningExecutions: a.EnableRestartRunningExecutions,

		IamRole:      a.IamRole,
		InstancePort: a.InstancePort,
		Name:         a.Name,

		Permissions:                    toClientApplicationPermissions(a.Permissions),
		PlatformHealthOnly:             a.PlatformHealthOnly,
		PlatformHealthOnlyShowOverride: a.PlatformHealthOnlyShowOverride,
		ProviderSettings:               a.toClientProviderSettings(a.ProviderSettings),

		RepoProjectKey: a.RepoProjectKey,
		RepoSlug:       a.RepoSlug,
		RepoType:       a.RepoType,
		TaskDefinition: a.TaskDefinition,
		TrafficGuards:  a.TrafficGuards,
	}
}

func fromClientApplication(a *client.Application) *application {
	return &application{
		Accounts:       fromClientAccounts(a.Accounts),
		CloudProviders: fromClientCloudProviders(a.CloudProviders),
		Email:          a.Email,

		EnableRestartRunningExecutions: a.EnableRestartRunningExecutions,

		IamRole:      a.IamRole,
		InstancePort: a.InstancePort,
		Name:         a.Name,

		Permissions:                    fromClientApplicationPermissions(a.Permissions),
		PlatformHealthOnly:             a.PlatformHealthOnly,
		PlatformHealthOnlyShowOverride: a.PlatformHealthOnlyShowOverride,
		ProviderSettings:               fromClientProviderSettings(a.ProviderSettings),

		RepoProjectKey: a.RepoProjectKey,
		RepoSlug:       a.RepoSlug,
		RepoType:       a.RepoType,
		TaskDefinition: a.TaskDefinition,
		TrafficGuards:  a.TrafficGuards,
	}
}

func fromClientAccounts(accounts string) []string {
	if len(accounts) == 0 {
		return []string{}
	}
	return strings.Split(accounts, ",")
}

func (a *application) setResourceData(d *schema.ResourceData) error {
	if err := d.Set("accounts", a.Accounts); err != nil {
		return err
	}
	if err := d.Set("cloud_providers", a.CloudProviders); err != nil {
		return err
	}
	if err := d.Set("email", a.Email); err != nil {
		return err
	}
	if err := d.Set("enable_restart_running_executions", a.EnableRestartRunningExecutions); err != nil {
		return err
	}
	if err := d.Set("instance_port", a.InstancePort); err != nil {
		return err
	}
	if err := d.Set("name", a.Name); err != nil {
		return err
	}
	if err := d.Set("permissions", a.Permissions); err != nil {
		return err
	}
	if err := d.Set("platform_health_only", a.PlatformHealthOnly); err != nil {
		return err
	}
	if err := d.Set("platform_health_only_show_override", a.PlatformHealthOnlyShowOverride); err != nil {
		return err
	}
	if err := d.Set("provider_settings", a.ProviderSettings); err != nil {
		return err
	}
	if err := d.Set("repo_project_key", a.RepoProjectKey); err != nil {
		return err
	}
	if err := d.Set("repo_slug", a.RepoSlug); err != nil {
		return err
	}
	return d.Set("repo_type", a.RepoType)
}
