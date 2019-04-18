package loqserver

type Config struct {
	ConfigurationPath string
}

func (c *Config) setupConfigurationVariable() {

}

func NewConfig() *Config {
	var config Config
	config.ConfigurationPath = "./config.ini"
	config.setupConfigurationVariable()
	return &config
}
