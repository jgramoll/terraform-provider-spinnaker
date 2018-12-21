package provider

// Trigger for Pipeline
type Trigger struct {
	ID           string
	Enabled      bool
	Job          string
	Master       string
	PropertyFile string `mapstructure:"property_file"`
	Type         string
}
