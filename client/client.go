package client

import (
  "bytes"
  "encoding/json"
  "net/http"
  "net/url"
)

type Config struct {
  Address   string
  CertPath  string
  KeyPath   string
}

type Client struct {
  Config  Config
  client  *http.Client
}

func NewClient(config Config) *Client {
  return &Client{Config: config, client: http.DefaultClient}
}

func (client *Client) Get(path string) (*http.Request, error) {
  return client.NewRequest("get", path)
}

func (client *Client) NewRequest(method string, path string) (*http.Request, error) {
  return client.NewRequestWithBody(method, path, nil)
}

func (client *Client) NewRequestWithBody(method string, path string, data map[string]string) (*http.Request, error) {
  reqUrl, urlErr := url.Parse(client.Config.Address + path)
  if urlErr != nil {
    return nil, urlErr
  }

  jsonValue, jsonErr := json.Marshal(data)
  if jsonErr != nil {
    return nil, jsonErr
  }

  return http.NewRequest(method, reqUrl.String(), bytes.NewBuffer(jsonValue))
}

func (client *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
  resp, err := client.client.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  // if err := validateResponse(resp); err != nil {
  //   return resp, err
  // }

  // err = decodeResponse(resp, v)
  return resp, err
}

// func decodeResponse(r *http.Response, v interface{}) error {
//   if v == nil {
//     return log.Errorf("nil interface provided to decodeResponse")
//   }

//   bodyBytes, _ := ioutil.ReadAll(r.Body)
//   bodyString := string(bodyBytes)
//   err := json.Unmarshal([]byte(bodyString), &v)
//   return err
// }