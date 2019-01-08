package provider

type moniker struct {
	App     string `mapstructure:"app"`
	Cluster string `mapstructure:"cluster"`
	Detail  string `mapstructure:"detail"`
	Stack   string `mapstructure:"stack"`
}
