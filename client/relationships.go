package client

// Relationships relationships
type Relationships struct {
	LoadBalancers  *[]string `json:"loadBalancers"`
	SecurityGroups *[]string `json:"securityGroups"`
}

// NewRelationships new relationships
func NewRelationships() *Relationships {
	return &Relationships{
		LoadBalancers:  &[]string{},
		SecurityGroups: &[]string{},
	}
}
