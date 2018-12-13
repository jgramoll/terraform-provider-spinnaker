package client

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestGetApplications(t *testing.T) {
	_, err := client.GetApplications()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateDeleteApplication(t *testing.T) {
	app := NewApplication()
	app.Name = fmt.Sprintf("MyTestApp%d", rand.Int())
	app.Email = client.Config.UserEmail
	err := client.CreateApplication(app)
	if err != nil {
		t.Fatal(err)
	}

	err = client.DeleteApplication(app)
	if err != nil {
		t.Fatal(err)
	}
}
