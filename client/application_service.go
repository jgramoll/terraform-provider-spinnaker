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

func (service *ApplicationService) getApplicationByName(name string) (*Application, error) {
	applications, err := service.GetApplications()
	if err != nil {
		return nil, err
	}
	if len(*applications) > 0 {
		for _, app := range *applications {
			if app.Name == name {
				return app, nil
			}
		}
	}
	return nil, fmt.Errorf("Could not find application with name \"%s\"", name)
}

// CreateApplication create an application
func (service *ApplicationService) CreateApplication(app *Application) error {
	if strings.Contains(app.Name, " ") {
		return ErrInvalidApplicationName
	}
	jobType := "createApplication"
	taskDescription := fmt.Sprintf("Create Application: %s", app.Name)
	return service.sendTask(app, jobType, taskDescription)
}

// DeleteApplication delete an application
func (service *ApplicationService) DeleteApplication(app *Application) error {
	jobType := "deleteApplication"
	taskDescription := fmt.Sprintf("Deleting Application: %s", app.Name)
	return service.sendTask(app, jobType, taskDescription)
}

func (service *ApplicationService) sendTask(app *Application, jobType string, taskDescription string) error {
	task := Task{
		Job: &[]*Job{
			&Job{
				Type:        jobType,
				Application: app,
				User:        service.Config.Auth.UserEmail,
			},
		},
		Application: app.Name,
		Description: taskDescription,
	}

	path := fmt.Sprintf("/applications/%s/tasks", app.Name)
	req, err := service.NewRequestWithBody("POST", path, task)
	if err != nil {
		return err
	}

	_, respErr := service.Do(req)
	return respErr
}
