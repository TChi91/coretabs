package config

//AppConfig structure
type AppConfig struct {
	FrontEnd FrontEndConfiguration `mapstructure:"frontend"`
	BackEnd  BackEndConfiguration  `mapstructure:"backend"`
}

// NewConfig Creates new instance of AppConfig
func NewConfig() AppConfig {
	config := AppConfig{}
	config.BackEnd.Port = 8000
	config.FrontEnd.Port = 8080

	return config
}
