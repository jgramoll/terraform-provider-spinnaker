package client

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
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

// GetApplicationByName return the application given a name
func (service *ApplicationService) GetApplicationByName(name string) (*Application, error) {
	path := fmt.Sprintf("/applications/%s", name)
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var attributes ApplicationAttributes
	_, err = service.DoWithResponse(req, &attributes)
	if err != nil {
		return nil, err
	}

	return attributes.Application, nil
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

// UpdateApplication update an application
func (service *ApplicationService) UpdateApplication(app *Application) error {
	jobType := "updateApplication"
	taskDescription := fmt.Sprintf("Updating Application: %s", app.Name)
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

	var taskResp TaskResponse
	_, err = service.DoWithResponse(req, taskResp)
	if err != nil {
		return err
	}

	req, err = service.NewRequest("GET", taskResp.Ref)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)

	go func() {
		var execution TaskExecution
		_, err = service.DoWithResponse(req, execution)
		if err != nil {
			log.Printf("[ERROR] Error on execute request to check task status. %s", err)
			done <- true
		}

		if execution.Status == "ERROR" {
			err = errors.New("Error on execute tasks")
			done <- true
		}

		if execution.Status == "SUCCEEDED" {
			done <- true
		}
	}()

	for {
		select {
		case <-done:
			return nil
		case <-ticker.C:
			continue
		}
	}

	return err
}
