package client

// ApplicationPermissions application permissions
type ApplicationPermissions struct {
	Execute []string `json:"EXECUTE"`
	Read    []string `json:"READ"`
	Write   []string `json:"WRITE"`
}

// NewApplicationPermissions new application permissions
func NewApplicationPermissions() *ApplicationPermissions {
	return &ApplicationPermissions{}
}
