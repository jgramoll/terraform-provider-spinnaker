package client

// Trigger for Pipeline
type Trigger struct {
	ID           string `json:"id"`
	Enabled      bool   `json:"enabled"`
	Job          string `json:"job"`
	Master       string `json:"master"`
	PropertyFile string `json:"propertyFile"`
	Type         string `json:"type"`
}
