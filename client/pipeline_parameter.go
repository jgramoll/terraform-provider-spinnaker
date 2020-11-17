package client

import (
	"errors"
)

// ErrParameterNotFound parameter not found
var ErrParameterNotFound = errors.New("could not find parameter")

// PipelineParameterOption option
type PipelineParameterOption struct {
	Value string `json:"value"`
}

// PipelineParameter parameter
type PipelineParameter struct {
	ID          string                      `json:"id"`
	Default     string                      `json:"default"`
	Description string                      `json:"description"`
	HasOptions  bool                        `json:"hasOptions"`
	Label       string                      `json:"label"`
	Name        string                      `json:"name"`
	Options     *[]*PipelineParameterOption `json:"options"`
	Required    bool                        `json:"required"`
}

// GetParameter by ID
func (p *Pipeline) GetParameter(parameterID string) (*PipelineParameter, error) {
	if p.ParameterConfig != nil {
		for _, parameter := range *p.ParameterConfig {
			if parameter.ID == parameterID {
				return parameter, nil
			}
		}
	}
	return nil, ErrParameterNotFound
}

// AppendParameter append parameter
func (p *Pipeline) AppendParameter(parameter *PipelineParameter) {
	if p.ParameterConfig == nil {
		p.ParameterConfig = &[]*PipelineParameter{}
	}
	config := append(*p.ParameterConfig, parameter)
	p.ParameterConfig = &config
}

// UpdateParameter in pipeline
func (p *Pipeline) UpdateParameter(parameter *PipelineParameter) error {
	if p.ParameterConfig != nil {
		for i, t := range *p.ParameterConfig {
			if t.ID == parameter.ID {
				(*p.ParameterConfig)[i] = parameter
				return nil
			}
		}
	}
	return ErrParameterNotFound
}

// DeleteParameter in pipeline
func (p *Pipeline) DeleteParameter(parameterID string) error {
	if p.ParameterConfig != nil {
		parameters := *p.ParameterConfig
		for i, t := range parameters {
			if t.ID == parameterID {
				parameters = append(parameters[:i], parameters[i+1:]...)
				p.ParameterConfig = &parameters
				return nil
			}
		}
	}
	return ErrParameterNotFound
}
