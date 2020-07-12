package config

//AppConfig structure
type AppConfig struct {
	FrontEnd FrontEndConfiguration `yaml:"frontend"`
	BackEnd  BackEndConfiguration  `yaml:"backend"`
}
