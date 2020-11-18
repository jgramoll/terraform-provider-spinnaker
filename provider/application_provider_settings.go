package provider

import (
	"strings"

	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type providerSettings struct {
	AWS *[]awsProviderSettings `mapstructure:"aws"`
}

type awsProviderSettings struct {
	UseAmiBlockDeviceMappings bool `mapstructure:"use_ami_block_device_mappings"`
}

func (a *application) toClientProviderSettings(settings *[]providerSettings) *client.ProviderSettings {
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

func fromClientCloudProviders(cloudProviders string) []string {
	if len(cloudProviders) == 0 {
		return []string{}
	}
	return strings.Split(cloudProviders, ",")
}

func fromClientProviderSettings(settings *client.ProviderSettings) *[]providerSettings {
	if settings == nil || settings.AWS == nil {
		return nil
	}

	return &[]providerSettings{
		{
			AWS: &[]awsProviderSettings{
				{UseAmiBlockDeviceMappings: settings.AWS.UseAmiBlockDeviceMappings},
			},
		},
	}
}
