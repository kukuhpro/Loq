package loqserver

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	ConfigurationPath string
}

func (c *Config) Get(key string) string {
	return os.Getenv(key)
}

func (c *Config) setValueConfiguration(array []string) bool {
	for _, val := range array {
		if !strings.Contains(val, "#") {
			r, _ := regexp.Compile("(.*)=(.*)")
			listString := r.FindStringSubmatch(val)
			if len(listString) >= 3 {
				os.Setenv(listString[1], listString[2])
			}
		}
	}
	return true
}

func (c *Config) setupConfigurationVariable() bool {
	file, err := os.Open(c.ConfigurationPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	arraySplitString := strings.Split(string(b), "\n")
	return c.setValueConfiguration(arraySplitString)
}

func NewConfig() *Config {
	var config Config
	config.ConfigurationPath = "./config.ini"
	config.setupConfigurationVariable()
	return &config
}
