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
	AWS *[]awsProviderSettings `mapstructure:"aws`
}

type permissions struct {
	Read    []string `mapstructure:"read"`
	Execute []string `mapstructure:"execute"`
	Write   []string `mapstructure:"write"`
}

// Application deploy application in application
type application struct {
	ID                             string              `mapstructure:"id"`
	Name                           string              `mapstructure:"name"`
	Email                          string              `mapstructure:"email"`
	RepoType                       string              `mapstructure:"repo_type"`
	RepoProjectKey                 string              `mapstructure:"repo_project_key"`
	RepoSlug                       string              `mapstructure:"repo_slug"`
	CloudProviders                 []string            `mapstructure:"cloud_providers"`
	Permissions                    *[]permissions      `mapstructure:"permissions"`
	ProviderSettings               *[]providerSettings `mapstructure:"provider_settings"`
	InstancePort                   int                 `mapstructure:"instance_port"`
	PlatformHealthOnly             bool                `mapstructure:"platform_health_only"`
	PlatformHealthOnlyShowOverride bool                `mapstructure:"platform_health_only_show_override"`
	EnableRestartRunningExecutions bool                `mapstructure:"enable_restart_running_executions"`
}

func (a *application) toClientApplication() *client.Application {
	return &client.Application{
		Name:                           a.Name,
		Email:                          a.Email,
		CloudProviders:                 strings.Join(a.CloudProviders, ","),
		RepoType:                       a.RepoType,
		RepoProjectKey:                 a.RepoProjectKey,
		RepoSlug:                       a.RepoSlug,
		Permissions:                    a.toPermissions(),
		ProviderSettings:               a.toClientProviderSettigns(a.ProviderSettings),
		InstancePort:                   a.InstancePort,
		PlatformHealthOnly:             a.PlatformHealthOnly,
		PlatformHealthOnlyShowOverride: a.PlatformHealthOnlyShowOverride,
		EnableRestartRunningExecutions: a.EnableRestartRunningExecutions,
	}
}

func (a *application) toPermissions() *client.Permissions {
	if a.Permissions == nil || len(*a.Permissions) == 0 {
		return nil
	}
	p := *a.Permissions
	return &client.Permissions{
		Read:    p[0].Read,
		Execute: p[0].Execute,
		Write:   p[0].Write,
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

func fromClientApplication(a *client.Application) *application {
	return &application{
		ID:                             a.Name,
		Name:                           a.Name,
		Email:                          a.Email,
		RepoType:                       a.RepoType,
		RepoProjectKey:                 a.RepoProjectKey,
		RepoSlug:                       a.RepoSlug,
		CloudProviders:                 strings.Split(a.CloudProviders, ","),
		Permissions:                    fromClientPermissions(a.Permissions),
		ProviderSettings:               fromClientProviderSettings(a.ProviderSettings),
		InstancePort:                   a.InstancePort,
		PlatformHealthOnly:             a.PlatformHealthOnly,
		PlatformHealthOnlyShowOverride: a.PlatformHealthOnlyShowOverride,
		EnableRestartRunningExecutions: a.EnableRestartRunningExecutions,
	}
}

func fromClientPermissions(p *client.Permissions) *[]permissions {
	if p == nil {
		return nil
	}
	return &[]permissions{{
		Read:    p.Read,
		Execute: p.Execute,
		Write:   p.Write,
	}}
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
	d.SetId(a.ID)
	err := d.Set("name", a.Name)
	if err != nil {
		return err
	}
	err = d.Set("email", a.Email)
	if err != nil {
		return err
	}
	err = d.Set("repo_type", a.RepoType)
	if err != nil {
		return err
	}
	err = d.Set("repo_project_key", a.RepoProjectKey)
	if err != nil {
		return err
	}
	err = d.Set("repo_slug", a.RepoSlug)
	if err != nil {
		return err
	}
	err = d.Set("instance_port", a.InstancePort)
	if err != nil {
		return err
	}
	err = d.Set("platform_health_only", a.PlatformHealthOnly)
	if err != nil {
		return err
	}
	err = d.Set("platform_health_only_show_override", a.PlatformHealthOnlyShowOverride)
	if err != nil {
		return err
	}
	err = d.Set("enable_restart_running_executions", a.EnableRestartRunningExecutions)
	if err != nil {
		return err
	}
	return d.Set("cloud_providers", a.CloudProviders)
}

// ApplicationFromResourceData get application from resource data
func applicationFromResourceData(application *client.Application, d *schema.ResourceData) {
	application.Name = d.Get("name").(string)
	application.Email = d.Get("email").(string)
	application.RepoType = d.Get("repo_type").(string)
	application.RepoProjectKey = d.Get("repo_project_key").(string)
	application.RepoSlug = d.Get("repo_slug").(string)
	application.CloudProviders = applicationCloudProvidersFromResourceData(d)
	application.Permissions = applicationPermissionsFromResourceData(d)
	application.ProviderSettings = applicationProviderSettingsFromResourceData(d)
	application.InstancePort = d.Get("instance_port").(int)
	application.PlatformHealthOnly = d.Get("platform_health_only").(bool)
	application.PlatformHealthOnlyShowOverride = d.Get("platform_health_only_show_override").(bool)
	application.EnableRestartRunningExecutions = d.Get("enable_restart_running_executions").(bool)
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

func applicationPermissionsFromResourceData(d *schema.ResourceData) *client.Permissions {
	PermissionsInterface, ok := d.GetOk("permissions")
	if !ok {
		return nil
	}

	clientPermissions := &client.Permissions{}
	for _, permissionsInterface := range PermissionsInterface.([]interface{}) {
		permissions := permissionsInterface.(map[string]interface{})
		clientPermissions.Read = typePermissionsFromResourceData(permissions["read"])
		clientPermissions.Write = typePermissionsFromResourceData(permissions["write"])
		clientPermissions.Execute = typePermissionsFromResourceData(permissions["execute"])
	}
	return clientPermissions
}

func typePermissionsFromResourceData(accesssListInterface interface{}) []string {
	listInterface, ok := accesssListInterface.([]interface{})
	if !ok || len(listInterface) == 0 {
		return nil
	}

	accessList := []string{}
	for _, accessInterface := range listInterface {
		accessList = append(accessList, accessInterface.(string))
	}
	return accessList
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
