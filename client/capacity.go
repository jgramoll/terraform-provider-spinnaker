package client

//Capacity capacity for cluster
type Capacity struct {
	Desired int `json:"desired"`
	Max     int `json:"max"`
	Min     int `json:"min"`
}
