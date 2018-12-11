package client

import (
  "io/ioutil"
  "testing"
)

var c Config
var client *Client

func init() {
  c = Config{
    Address: "http://example.com",
    CertPath: "#certPath",
    KeyPath: "KeyPath"}
  client = NewClient(c)
}

func TestClientNewRequest(t *testing.T) {
  req, err := client.NewRequest("get", "/test/path")
  if (err != nil) {
    t.Fatal(err)
  }
  expectedUrl := "http://example.com/test/path"
  if (req.URL.String() != expectedUrl) {
    t.Fatalf("request url should be %#v, not %#v", expectedUrl, req.URL.String())
  }
}

func TestClientNewRequestWithBody(t *testing.T) {
  body := map[string]string {
    "field": "#value",
  }
  req, err := client.NewRequestWithBody("get", "/test/path", body)
  if (err != nil) {
    t.Fatal(err)
  }
  byteBody, bodyErr := ioutil.ReadAll(req.Body)
  if (bodyErr != nil) {
    t.Fatal(bodyErr)
  }

  actualBody := string(byteBody)
  expectedBody := "{\"field\":\"#value\"}"
  if (actualBody != expectedBody) {
    t.Fatalf("request body should be %#v, not %#v", actualBody, req.URL.String())
  }
}


func TestClientDo(t *testing.T) {
  req, err := client.NewRequest("get", "/test/path")
  if (err != nil) {
    t.Fatal(err)
  }
  t.Log(req)

  // TODO this actually sends...
  // resp, err := client.Do(req, nil)
  // if (err != nil) {
  //   t.Fatal(err)
  // }
  // TODO what can we assert?
}

