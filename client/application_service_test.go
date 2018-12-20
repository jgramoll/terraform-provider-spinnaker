package client

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

var applicationService *ApplicationService

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	applicationService = &ApplicationService{newTestClient()}
}

func TestCreateDeleteApplication(t *testing.T) {
	expectedName := fmt.Sprintf("mytestapp%d", rand.Int())
	app := NewApplication()
	app.Name = expectedName
	app.Email = applicationService.Config.UserEmail
	err := applicationService.CreateApplication(app)
	if err != nil {
		t.Fatal(err)
	}

	var savedApp *Application
	savedApp, err = applicationService.getApplicationByName(expectedName)
	if err != nil {
		t.Fatal(err)
	}
	err = savedApp.equals(app)
	if err != nil {
		t.Fatal(err)
	}

	err = applicationService.DeleteApplication(app)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApplicationNameWithSpace(t *testing.T) {
	expectedName := fmt.Sprintf("my test app %d", rand.Int())
	app := NewApplication()
	app.Name = expectedName
	err := applicationService.CreateApplication(app)
	if err == nil {
		t.Fatal("Should not allow spaces in application name")
	}
	if err != ErrInvalidApplicationName {
		t.Fatal(err)
	}
}

func TestApplicationCleanup(t *testing.T) {
	apps, err := applicationService.GetApplications()
	if err != nil {
		t.Fatal(err)
	}

	for _, app := range *apps {
		if strings.Contains(app.Name, "mytestapp") {
			applicationService.DeleteApplication(app)
		}
	}
}

// GetApplications get all applications
func (service *ApplicationService) getApplicationByName(name string) (*Application, error) {
	var applications *[]*Application
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

func (application *Application) equals(expectedApplication *Application) error {
	if !reflect.DeepEqual(application, expectedApplication) {
		return fmt.Errorf("Application %v does not match %v", application, expectedApplication)
	}
	return nil
}
