// +build integration

package client

import (
	"fmt"
	"math/rand"
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
	appName := fmt.Sprintf("mytestapp%d", rand.Int())
	app := NewApplication()
	app.Name = appName
	app.Email = applicationService.Config.Auth.UserEmail
	err := applicationService.CreateApplication(app)
	if err != nil {
		t.Fatal(err)
	}

	app, err = applicationService.GetApplicationByNameWithRetries(appName)
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
		if strings.Contains(app.Email, "@spin.com") {
			applicationService.DeleteApplication(app)
		}
		if strings.Contains(app.Name, "mytestapp") {
			applicationService.DeleteApplication(app)
		}
	}
}
