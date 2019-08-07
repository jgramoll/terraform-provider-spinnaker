package client

type Relationships struct {
	LoadBalancers  *[]string `json:"loadBalancers"`
	SecurityGroups *[]string `json:"securityGroups"`
}

func NewRelationships() *Relationships {
	return &Relationships{
		LoadBalancers:  &[]string{},
		SecurityGroups: &[]string{},
	}
}
