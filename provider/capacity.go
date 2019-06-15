package provider

type capacity struct {
	Desired string `mapstructure:"desired"`
	Max     string `mapstructure:"max"`
	Min     string `mapstructure:"min"`
}
