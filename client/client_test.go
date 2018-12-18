package client

import (
	"fmt"
	"io/ioutil"
	"log"
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

func TestClientDo(t *testing.T) {
	req, err := client.NewRequest("get", testPath)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(req)

	// TODO this actually sends... what can we assert?
	// https://golang.org/pkg/net/http/httptest/#example_ResponseRecorder
	// resp, err := client.Do(req, nil)
	// if (err != nil) {
	//   t.Fatal(err)
	// }
}

func newTestClient() *Client {
	usr, err := user.Current()
	if err != nil {
		log.Println("[Error] unable to get current user: ", err)
	}

	c := Config{
		Address:   "https://api.spinnaker.inseng.net",
		CertPath:  usr.HomeDir + "/.spin/client.crt",
		KeyPath:   usr.HomeDir + "/.spin/client.key",
		UserEmail: fmt.Sprintf("%s@instructure.com", usr.Username),
	}
	return NewClient(c)
}
