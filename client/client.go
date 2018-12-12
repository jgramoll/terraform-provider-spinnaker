package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Config struct {
	Address  string
	CertPath string
	KeyPath  string
}

type Client struct {
	Config Config
	client *http.Client
}

func NewClient(config Config) *Client {
	cert, err := tls.LoadX509KeyPair(config.CertPath, config.KeyPath)
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	c := &http.Client{Transport: transport}

	return &Client{Config: config, client: c}
}

func (client *Client) Get(path string) (*http.Request, error) {
	return client.NewRequest("get", path)
}

func (client *Client) NewRequest(method string, path string) (*http.Request, error) {
	return client.NewRequestWithBody(method, path, nil)
}

func (client *Client) NewRequestWithBody(method string, path string, data map[string]interface{}) (*http.Request, error) {
	reqUrl, urlErr := url.Parse(client.Config.Address + path)
	if urlErr != nil {
		return nil, urlErr
	}

	jsonValue, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		return nil, jsonErr
	}

	req, err := http.NewRequest(method, reqUrl.String(), bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	return req, nil
}

func (client *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := validateResponse(resp); err != nil {
		return resp, err
	}

	err = decodeResponse(resp, v)
	return resp, err

}

func decodeResponse(r *http.Response, v interface{}) error {
	if v == nil {
		return fmt.Errorf("nil interface provided to decodeResponse")
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	log.Printf("[INFO] Got response body %s\n", bodyString)

	err := json.Unmarshal([]byte(bodyString), &v)
	return err
}

func validateResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	log.Printf("[INFO] Error response body %s\n", bodyString)
	error := &SpinnakerError{}
	err := json.Unmarshal([]byte(bodyString), &error)
	if err != nil {
		return err
	}

	return error
}
