package lib

import "testing"

type ConfigTestSuite struct {
	config *Config
}

var configTestSuite *ConfigTestSuite

func init() {
	configTestSuite = &ConfigTestSuite{}
	configTestSuite.config = NewConfig()
}

func TestSetupConfiguration(t *testing.T) {
	config := configTestSuite.config
	resultValue := config.setupConfigurationVariable()
	if !resultValue {
		t.Errorf("Result value must be True")
	}
}

var tableDrivenGetConfiguration = map[string]string{
	"produce_log":       "folder",
	"path_folder_local": "log",
}

func TestGetConfiguration(t *testing.T) {
	config := configTestSuite.config
	for k, v := range tableDrivenGetConfiguration {
		result := config.Get(k)
		if result != v {
			t.Errorf("You have actual result '" + result + "' and expected result '" + v + "'")
		}
	}
}
