package client

import (
	"errors"
	"fmt"
	"strings"
)

// ErrInvalidApplicationName invalid application name
var ErrInvalidApplicationName = errors.New("Invalid application name")

// ApplicationService for interacting with spinnaker applications
type ApplicationService struct {
	*Client
}

// GetApplications get all applications
func (service *ApplicationService) GetApplications() (*[]*Application, error) {
	path := "/applications"
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var apps []*Application
	_, respErr := service.DoWithResponse(req, &apps)
	if respErr != nil {
		return nil, respErr
	}

	return &apps, nil
}

// CreateApplication create an application
func (service *ApplicationService) CreateApplication(app *Application) error {
	if strings.Contains(app.Name, " ") {
		return ErrInvalidApplicationName
	}

	task := Task{
		Job: &[]*Job{
			&Job{
				Type:        "createApplication",
				Application: app,
				User:        service.Config.UserEmail,
			},
		},
		Application: app.Name,
		Description: fmt.Sprintf("Create Application: %s", app.Name),
	}

	path := fmt.Sprintf("/applications/%s/tasks", app.Name)
	req, err := service.NewRequestWithBody("POST", path, task)
	if err != nil {
		return err
	}

	_, respErr := service.Do(req)
	return respErr
}

// DeleteApplication delete an application
func (service *ApplicationService) DeleteApplication(app *Application) error {
	task := Task{
		Job: &[]*Job{
			&Job{
				Type:        "deleteApplication",
				Application: app,
				User:        service.Config.UserEmail,
			},
		},
		Application: app.Name,
		Description: fmt.Sprintf("Deleting Application: %s", app.Name),
	}

	path := fmt.Sprintf("/applications/%s/tasks", app.Name)
	req, err := service.NewRequestWithBody("POST", path, task)
	if err != nil {
		return err
	}

	_, respErr := service.Do(req)
	return respErr
}
