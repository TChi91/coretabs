package config

//FrontEndConfiguration structure
type FrontEndConfiguration struct {
	PackageManager string `mapstructure:"package-manager"`
	Port           int    `mapstructure:"port"`
	Server         string `mapstructure:"server"`
}
