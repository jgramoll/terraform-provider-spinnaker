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
	CloudProviders                 string            `json:"cloudProviders"`
	InstancePort                   int               `json:"instancePort"`
	ProviderSettings               *ProviderSettings `json:"providerSettings"`
	Name                           string            `json:"name"`
	Email                          string            `json:"email"`
	RepoType                       string            `json:"repoType"`
	RepoProjectKey                 string            `json:"repoProjectKey"`
	RepoSlug                       string            `json:"repoSlug"`
	PlatformHealthOnly             bool              `json:"platformHealthOnly"`
	PlatformHealthOnlyShowOverride bool              `json:"platformHealthOnlyShowOverride"`
	EnableRestartRunningExecutions bool              `json:"enableRestartRunningExecutions"`
}

// ApplicationAttributes mapping for `application/{appName}`  endpoint
type ApplicationAttributes struct {
	Application *Application `json:"attributes"`
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
