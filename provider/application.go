package provider

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type awsProviderSettings struct {
	UseAmiBlockDeviceMappings bool `mapstructure:"use_ami_block_device_mappings"`
}

type providerSettings struct {
	AWS *[]awsProviderSettings `mapstructure:"aws"`
}

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

	PlatformHealthOnly             bool                `mapstructure:"platform_health_only"`
	PlatformHealthOnlyShowOverride bool                `mapstructure:"platform_health_only_show_override"`
	ProviderSettings               *[]providerSettings `mapstructure:"provider_settings"`

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

		PlatformHealthOnly:             a.PlatformHealthOnly,
		PlatformHealthOnlyShowOverride: a.PlatformHealthOnlyShowOverride,
		ProviderSettings:               a.toClientProviderSettigns(a.ProviderSettings),

		RepoProjectKey: a.RepoProjectKey,
		RepoSlug:       a.RepoSlug,
		RepoType:       a.RepoType,
		TaskDefinition: a.TaskDefinition,
		TrafficGuards:  a.TrafficGuards,
	}
}

func (a *application) toClientProviderSettigns(settings *[]providerSettings) *client.ProviderSettings {
	if settings != nil || len(*settings) > 0 {
		for _, setting := range *settings {
			if setting.AWS != nil && len(*setting.AWS) > 0 {
				for _, aws := range *setting.AWS {
					return &client.ProviderSettings{
						AWS: &client.AwsProviderSettings{
							UseAmiBlockDeviceMappings: aws.UseAmiBlockDeviceMappings,
						},
					}
				}
			}
		}
	}

	return nil
}

func fromClientAccounts(accounts string) []string {
	if len(accounts) == 0 {
		return []string{}
	}
	return strings.Split(accounts, ",")
}

func fromClientCloudProviders(cloudProviders string) []string {
	if len(cloudProviders) == 0 {
		return []string{}
	}
	return strings.Split(cloudProviders, ",")
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

func fromClientProviderSettings(settings *client.ProviderSettings) *[]providerSettings {
	if settings == nil || settings.AWS == nil {
		return nil
	}

	return &[]providerSettings{
		providerSettings{
			AWS: &[]awsProviderSettings{
				{UseAmiBlockDeviceMappings: settings.AWS.UseAmiBlockDeviceMappings},
			},
		},
	}
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

// ApplicationFromResourceData get application from resource data
func applicationFromResourceData(application *client.Application, d *schema.ResourceData) {
	application.Accounts = applicationAccountsFromResourceData(d)
	application.CloudProviders = applicationCloudProvidersFromResourceData(d)
	application.Email = d.Get("email").(string)
	application.EnableRestartRunningExecutions = d.Get("enable_restart_running_executions").(bool)
	application.InstancePort = d.Get("instance_port").(int)
	application.Name = d.Get("name").(string)
	application.PlatformHealthOnly = d.Get("platform_health_only").(bool)
	application.PlatformHealthOnlyShowOverride = d.Get("platform_health_only_show_override").(bool)
	application.ProviderSettings = applicationProviderSettingsFromResourceData(d)
	application.RepoProjectKey = d.Get("repo_project_key").(string)
	application.RepoSlug = d.Get("repo_slug").(string)
	application.RepoType = d.Get("repo_type").(string)
}

func applicationAccountsFromResourceData(d *schema.ResourceData) string {
	accountsInterface, ok := d.GetOk("accounts")
	if !ok {
		return ""
	}

	accounts := []string{}
	for _, a := range accountsInterface.([]interface{}) {
		accounts = append(accounts, a.(string))
	}
	return strings.Join(accounts, ",")
}

func applicationCloudProvidersFromResourceData(d *schema.ResourceData) string {
	cloudProvidersInterface, ok := d.GetOk("cloud_providers")
	if !ok {
		return ""
	}

	cloudProviders := []string{}
	for _, cloudProvider := range cloudProvidersInterface.([]interface{}) {
		cloudProviders = append(cloudProviders, cloudProvider.(string))
	}
	return strings.Join(cloudProviders, ",")
}

func applicationProviderSettingsFromResourceData(d *schema.ResourceData) *client.ProviderSettings {
	providerSettingsInterface, ok := d.GetOk("provider_settings")
	if !ok {
		return nil
	}

	clientProviderSettings := &client.ProviderSettings{}
	for _, providerSettingInterface := range providerSettingsInterface.([]interface{}) {
		providerSetting := providerSettingInterface.(map[string]interface{})
		clientProviderSettings.AWS = awsProviderSettingsFromResourceData(providerSetting)
	}
	return clientProviderSettings
}

func awsProviderSettingsFromResourceData(providerSetting map[string]interface{}) *client.AwsProviderSettings {
	for _, awsInterface := range providerSetting["aws"].([]interface{}) {
		aws := awsInterface.(map[string]interface{})
		return &client.AwsProviderSettings{
			UseAmiBlockDeviceMappings: aws["use_ami_block_device_mappings"].(bool),
		}
	}

	return nil
}
