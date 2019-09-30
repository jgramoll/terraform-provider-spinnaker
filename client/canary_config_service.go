package client

import (
	"errors"
	"fmt"
)

var ErrCanaryConfigNotFound = errors.New("Could not find canary config")

type CreateCanaryConfigResponse struct {
	CanaryConfigId string `json:"canaryConfigId"`
}

type CanaryConfigService struct {
	*Client
}

func (service *CanaryConfigService) GetCanaryConfigs() (*[]*CanaryConfig, error) {
	path := "/v2/canaryConfig"
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var configs []*CanaryConfig
	_, respErr := service.DoWithResponse(req, &configs)
	if respErr != nil {
		return nil, respErr
	}

	return &configs, nil
}

func (service *CanaryConfigService) GetCanaryConfig(id string) (*CanaryConfig, error) {
	path := fmt.Sprintf("/v2/canaryConfig/%s", id)
	req, err := service.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	config := &CanaryConfig{}
	_, err = service.DoWithResponse(req, config)
	return config, err
}

// CreateApplication create an application
func (service *CanaryConfigService) CreateCanaryConfig(config *CanaryConfig) (configId string, err error) {
	path := "/v2/canaryConfig"
	req, err := service.NewRequestWithBody("POST", path, config)
	if err != nil {
		return "", err
	}

	response := &CreateCanaryConfigResponse{}
	_, err = service.DoWithResponse(req, response)
	return response.CanaryConfigId, err
}

func (service *CanaryConfigService) UpdateCanaryConfig(config *CanaryConfig) error {
	path := fmt.Sprintf("/v2/canaryConfig/%s", config.Id)
	req, err := service.NewRequestWithBody("PUT", path, config)
	if err != nil {
		return err
	}

	_, err = service.Do(req)
	return err
}

func (service *CanaryConfigService) DeleteCanaryConfig(configId string) error {
	path := fmt.Sprintf("/v2/canaryConfig/%s", configId)
	req, err := service.NewRequest("DELETE", path)
	if err != nil {
		return err
	}

	_, err = service.Do(req)
	return err
}
