package config

//BackEndConfiguration structure
type BackEndConfiguration struct {
	Port   int    `mapstructure:"port"`
	Server string `mapstructure:"server"`
}
