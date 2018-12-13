package client

import (
	"fmt"
)

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

// GetApplications get all applications
func (client *Client) GetApplications() (*[]*Application, error) {
	path := "/applications"
	req, err := client.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var apps []*Application
	_, respErr := client.DoWithResponse(req, &apps)
	if respErr != nil {
		return nil, respErr
	}

	return &apps, nil
}

// CreateApplication create an application
func (client *Client) CreateApplication(app *Application) error {
	// TODO no spaces in name lint
	task := Task{
		Job: &[]*Job{
			&Job{
				Type:        "createApplication",
				Application: app,
				User:        client.Config.UserEmail,
			},
		},
		Application: app.Name,
		Description: fmt.Sprintf("Create Application: %s", app.Name),
	}

	path := fmt.Sprintf("/applications/%s/tasks", app.Name)
	req, err := client.NewRequestWithBody("POST", path, task)
	if err != nil {
		return err
	}

	_, respErr := client.Do(req)
	return respErr
}

// DeleteApplication delete an application
func (client *Client) DeleteApplication(app *Application) error {
	task := Task{
		Job: &[]*Job{
			&Job{
				Type:        "deleteApplication",
				Application: app,
				User:        client.Config.UserEmail,
			},
		},
		Application: app.Name,
		Description: fmt.Sprintf("Deleting Application: %s", app.Name),
	}

	path := fmt.Sprintf("/applications/%s/tasks", app.Name)
	req, err := client.NewRequestWithBody("POST", path, task)
	if err != nil {
		return err
	}

	_, respErr := client.Do(req)
	return respErr
}
