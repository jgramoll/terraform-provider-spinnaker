package client

// AwsProviderSettings Settings for AWS Provider
type AwsProviderSettings struct {
	UseAmiBlockDeviceMappings bool `json:"useAmiBlockDeviceMappings"`
}

// ProviderSettings Settings for Provider
type ProviderSettings struct {
	AWS *AwsProviderSettings `json:"aws"`
}

// Application Settings for Application
type Application struct {
	CloudProviders   string            `json:"cloudProviders"`
	InstancePort     int               `json:"instancePort"`
	ProviderSettings *ProviderSettings `json:"providerSettings"`
	Name             string            `json:"name"`
	Email            string            `json:"email"`
}

// NewAwsProviderSettings return Aws provider settings with default values
func NewAwsProviderSettings() *AwsProviderSettings {
	return &AwsProviderSettings{
		UseAmiBlockDeviceMappings: false,
	}
}

// NewApplication return Application object with default values
func NewApplication() *Application {
	return &Application{
		InstancePort: 80,
		ProviderSettings: &ProviderSettings{
			AWS: NewAwsProviderSettings(),
		},
	}
}
