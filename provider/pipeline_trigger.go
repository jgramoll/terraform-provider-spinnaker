package provider

// Trigger for Pipeline
type Trigger struct {
	ID           string
	Enabled      bool
	Job          string
	Master       string
	PropertyFile string `mapstructure:"property_file"`
	RunAsUser    string `mapstructure:"run_as_user"`
	Type         string
}
