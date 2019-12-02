package client

type ApplicationPermissions struct {
	Execute []string `json:"EXECUTE"`
	Read    []string `json:"READ"`
	Write   []string `json:"WRITE"`
}

func NewApplicationPermissions() *ApplicationPermissions {
	return &ApplicationPermissions{}
}
