package config

//AppConfig structure
type AppConfig struct {
	FrontEnd FrontEndConfiguration `yaml:"frontend"`
	BackEnd  BackEndConfiguration  `yaml:"backend"`
}

// NewConfig Creates new instance of AppConfig
func NewConfig() AppConfig {
	config := AppConfig{}
	config.BackEnd.Port = 8000
	config.FrontEnd.Port = 8080

	return config
}
