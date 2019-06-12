package client

//Capacity capacity for cluster
type Capacity struct {
	Desired string `json:"desired"`
	Max     string `json:"max"`
	Min     string `json:"min"`
}
