package client

type PipelineParameterOption struct {
	Value string `json:"value"`
}

type PipelineParameter struct {
	Default     string                      `json:"default"`
	Description string                      `json:"description"`
	HasOptions  bool                        `json:"hasOptions"`
	Label       string                      `json:"label"`
	Name        string                      `json:"name"`
	Options     *[]*PipelineParameterOption `json:"options"`
	Required    bool                        `json:"required"`
}
