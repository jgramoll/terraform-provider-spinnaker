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

// ApplicationAttributes mapping for `application/{appName}`  endpoint
type ApplicationAttributes struct {
	Application *Application `json:"attributes"`
}

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

// GetApplicationByNameWithRetries return the application given a name and retries
func (service *ApplicationService) GetApplicationByNameWithRetries(name string) (*Application, error) {

	// Execute every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	appChan := make(chan *Application)
	errChan := make(chan error)

	f := func() {
		app, err := service.GetApplicationByName(name)
		if err != nil {
			spinnakerError, ok := err.(*SpinnakerError)
			log.Printf("[DEBUG] GetApplicationByName error: %s\n", err.Error())
			if !ok || spinnakerError.Status != 403 {
				errChan <- fmt.Errorf("Error finding app %v.\n %v", name, err)
			}
		}
		if app != nil {
			appChan <- app
			log.Printf("[DEBUG] Found application: %s\n", name)
		}
	}

	go func() {
		time.Sleep(10 * time.Minute)
		errChan <- fmt.Errorf("Execution timeout on finding app %s", name)
	}()

	for {
		go f()

		select {
		case app := <-appChan:
			return app, nil
		case err := <-errChan:
			return nil, err
		case <-ticker.C:
			continue
		}
	}
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
			{
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
	_, err = service.DoWithResponse(req, &taskResp)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Checking job %s task %s execution\n", jobType, taskResp.Ref)
	req, err = service.NewRequest("GET", taskResp.Ref)
	if err != nil {
		return err
	}

	// Execute every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	done := make(chan bool)
	errChan := make(chan error)

	f := func() {
		var execution TaskExecution
		_, err := service.DoWithResponse(req, &execution)
		if err != nil {
			log.Printf("[ERROR] Error on execute request to check task status: %s\n", err)
			return
		}

		log.Printf("[DEBUG] Task %s current status %s\n", jobType, execution.Status)
		if execution.Status == "ERROR" || execution.Status == "TERMINAL" {
			errChan <- fmt.Errorf("Error on execute job %s task id %s", jobType, taskResp.Ref)
		}

		if execution.Status == "SUCCEEDED" {
			log.Printf("[DEBUG] Task %s finished with SUCCEEDED\n", jobType)
			done <- true
		}
	}

	// Wait for 5 minutes before failing execution
	go func() {
		time.Sleep(5 * time.Minute)
		errChan <- fmt.Errorf("Execution timeout on %s task id %s", jobType, taskResp.Ref)
	}()

	for {
		go f()

		select {
		case <-done:
			return nil
		case err := <-errChan:
			return err
		case <-ticker.C:
			continue
		}
	}
}
