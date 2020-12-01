package client

import (
	"errors"
	"fmt"
)

// ErrCanaryConfigNotFound config not found
var ErrCanaryConfigNotFound = errors.New("Could not find canary config")

// CreateCanaryConfigResponse canary config response
type CreateCanaryConfigResponse struct {
	CanaryConfigID string `json:"canaryConfigId"`
}

// CanaryConfigService config service
type CanaryConfigService struct {
	*Client
}

// GetCanaryConfigs get configs
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

// GetCanaryConfig get config by id
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

// CreateCanaryConfig create an canary config
func (service *CanaryConfigService) CreateCanaryConfig(config *CanaryConfig) (configID string, err error) {
	path := "/v2/canaryConfig"
	req, err := service.NewRequestWithBody("POST", path, config)
	if err != nil {
		return "", err
	}

	response := &CreateCanaryConfigResponse{}
	_, err = service.DoWithResponse(req, response)
	return response.CanaryConfigID, err
}

// UpdateCanaryConfig update config
func (service *CanaryConfigService) UpdateCanaryConfig(config *CanaryConfig) error {
	path := fmt.Sprintf("/v2/canaryConfig/%s", config.ID)
	req, err := service.NewRequestWithBody("PUT", path, config)
	if err != nil {
		return err
	}

	_, err = service.Do(req)
	return err
}

// DeleteCanaryConfig deleete config
func (service *CanaryConfigService) DeleteCanaryConfig(configID string) error {
	path := fmt.Sprintf("/v2/canaryConfig/%s", configID)
	req, err := service.NewRequest("DELETE", path)
	if err != nil {
		return err
	}

	_, err = service.Do(req)
	return err
}
