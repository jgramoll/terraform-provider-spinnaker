package provider

type capacity struct {
	Desired           int    `mapstructure:"desired"`
	DesiredExpression string `mapstructure:"desired_expression"`
	Max               int    `mapstructure:"max"`
	MaxExpression     string `mapstructure:"max_expression"`
	Min               int    `mapstructure:"min"`
	MinExpression     string `mapstructure:"min_expression"`
}
