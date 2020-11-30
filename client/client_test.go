// +build integration

package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"testing"
)

var client *Client
var testPath string

func init() {
	testPath = "/test/path"
	client = newTestClient()
}

func TestClientNewRequest(t *testing.T) {
	req, err := client.NewRequest("get", testPath)
	if err != nil {
		t.Fatal(err)
	}
	expectedURL := client.Config.Address + testPath
	if req.URL.String() != expectedURL {
		t.Fatalf("request url should be %#v, not %#v", expectedURL, req.URL.String())
	}
}

func TestClientNewRequestWithBody(t *testing.T) {
	body := map[string]interface{}{
		"field": "#value",
	}
	req, err := client.NewRequestWithBody("get", testPath, body)
	if err != nil {
		t.Fatal(err)
	}
	byteBody, bodyErr := ioutil.ReadAll(req.Body)
	if bodyErr != nil {
		t.Fatal(bodyErr)
	}

	actualBody := string(byteBody)
	expectedBody := "{\"field\":\"#value\"}"
	if actualBody != expectedBody {
		t.Fatalf("request body should be %#v, not %#v", actualBody, req.URL.String())
	}
}

func TestClientErrorResponse(t *testing.T) {
	req, err := client.NewRequest("GET", testPath)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Do(req)
	if err == nil {
		t.Fatal("should fail")
	}
	spinnakerError, ok := err.(*SpinnakerError)
	if !ok {
		t.Fatalf("expected spinnaker error, got %v", err)
	}
	if spinnakerError.Status != 404 {
		t.Fatalf("should return 404, not %v", spinnakerError.Status)
	}
}

func newTestClient() *Client {
	usr, err := user.Current()
	if err != nil {
		log.Printf("[ERROR] unable to get current user: %s\n", err)
	}

	address := os.Getenv("SPINNAKER_ADDRESS")
	if address == "" {
		log.Println("[ERROR] SPINNAKER_ADDRESS not defined")
	}
	certPath := os.Getenv("SPINNAKER_CERT")
	if certPath == "" {
		log.Println("[ERROR] SPINNAKER_CERT not defined")
	}
	keyPath := os.Getenv("SPINNAKER_KEY")
	if keyPath == "" {
		log.Println("[ERROR] SPINNAKER_KEY not defined")
	}

	c := NewConfig()
	c.Address = address
	c.Auth.CertPath = certPath
	c.Auth.KeyPath = keyPath
	c.Auth.UserEmail = fmt.Sprintf("%s", usr.Username)
	client, err := NewClient(c)
	if err != nil {
		log.Printf("[ERROR] Creating client: %s\n", err)
	}
	return client
}
