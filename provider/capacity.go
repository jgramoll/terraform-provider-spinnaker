package provider

type capacity struct {
	Desired int `mapstructure:"desired"`
	Max     int `mapstructure:"max"`
	Min     int `mapstructure:"min"`
}
