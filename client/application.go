package client

// AwsProviderSettings Settings for AWS Provider
type AwsProviderSettings struct {
	UseAmiBlockDeviceMappings bool `json:"useAmiBlockDeviceMappings"`
}

// ProviderSettings Settings for Provider
type ProviderSettings struct {
	AWS *AwsProviderSettings `json:"aws"`
}

// DataSources data sources for application
type DataSources struct {
	Disabled *[]string `json:"disabled"`
	Enabled  *[]string `json:"enabled"`
}

// Application Settings for Application
type Application struct {
	Accounts       string       `json:"accounts"`
	CloudProviders   string            `json:"cloudProviders"`
	CreateTs       string       `json:"createTs"`
	DataSources    *DataSources `json:"dataSources"`
	DesiredCount   string       `json:"desiredCount"`
	Email          string       `json:"email"`

	EnableRestartRunningExecutions bool `json:"enableRestartRunningExecutions"`

	IamRole        string `json:"iamRole"`
	InstancePort     int               `json:"instancePort"`
	LastModifiedBy string `json:"lastModifiedBy"`
	Name           string `json:"name"`

	Permissions                    *ApplicationPermissions `json:"permissions"`
	PlatformHealthOnly             bool                    `json:"platformHealthOnly"`
	PlatformHealthOnlyShowOverride bool                    `json:"platformHealthOnlyShowOverride"`
	ProviderSettings *ProviderSettings `json:"providerSettings"`

	RepoProjectKey string   `json:"repoProjectKey"`
	RepoSlug       string   `json:"repoSlug"`
	RepoType       string   `json:"repoType"`
	TaskDefinition string   `json:"taskDefinition"`
	TrafficGuards  []string `json:"trafficGuards"`
	UpdateTs       string   `json:"updateTs"`
	User           string   `json:"user"`
	RepoType                       string            `json:"repoType"`
	RepoProjectKey                 string            `json:"repoProjectKey"`
	RepoSlug                       string            `json:"repoSlug"`
	Permissions                    *Permissions      `json:"permissions"`
	PlatformHealthOnly             bool              `json:"platformHealthOnly"`
	PlatformHealthOnlyShowOverride bool              `json:"platformHealthOnlyShowOverride"`
	EnableRestartRunningExecutions bool              `json:"enableRestartRunningExecutions"`
}

// ApplicationAttributes mapping for `application/{appName}`  endpoint
type ApplicationAttributes struct {
	Application *Application `json:"attributes"`
}

// Permissions define the authorized groups access
type Permissions struct {
	Read    []string `json:"READ"`
	Execute []string `json:"EXECUTE"`
	Write   []string `json:"WRITE"`
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
