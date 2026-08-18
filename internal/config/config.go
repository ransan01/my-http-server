package config

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

const ENV_VAR_PREFIX = "MY_API_SERVER_"

var CONFIG_NAMES = map[string]int{
	"Address":  0,
	"Port":     1,
	"CertFile": 2,
	"KeyFile":  3,
}

type Json struct {
	Name  string
	Value string
}

type EnvVar struct {
	Name  string
	Value string
}

type ConfigurationTypes struct {
	Json   Json
	EnvVar EnvVar
}

type Configuration struct {
	Types ConfigurationTypes
}

type Config struct {
	Configurations []Configuration
}

var config = &Config{
	Configurations: []Configuration{},
}

func init() {
	for name, _ := range CONFIG_NAMES {
		config.Configurations = append(config.Configurations, Configuration{
			Types: ConfigurationTypes{
				Json: Json{
					Name:  name,
					Value: "",
				},
				EnvVar: EnvVar{
					Name:  ENV_VAR_PREFIX + strings.ToUpper(name),
					Value: "",
				},
			},
		})
	}
}

func loadConfigFromEnv() {
	for i, configuration := range config.Configurations {
		config.Configurations[i].Types.EnvVar.Value = os.Getenv(configuration.Types.EnvVar.Name)
	}
}

func loadConfigFromJsonFile(filename string) (*Config, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	config := &Config{}
	err = json.NewDecoder(jsonFile).Decode(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func LoadConfig() error {
	loadConfigFromEnv()
	jsonConfig, err := loadConfigFromJsonFile("config.json")
	if err != nil {
		return err
	}
	// Merge the environment variables with the json config
	// The environment variables take precedence over the json config
	// If the environment variable is set, then json type config gets the same value as the environment variable
	// Each configuration's Json.Value and EnvVar.Value should be the same
	for _, jsonConfiguration := range jsonConfig.Configurations {
		for i, configuration := range config.Configurations {
			if configuration.Types.Json.Name == jsonConfiguration.Types.Json.Name {
				if configuration.Types.EnvVar.Value != "" {
					// Env var takes precedence — keep EnvVar.Value, sync Json.Value to it
					config.Configurations[i].Types.Json.Value = configuration.Types.EnvVar.Value
				} else {
					// No env var — use json file value, sync EnvVar.Value too
					config.Configurations[i].Types.Json.Value = jsonConfiguration.Types.Json.Value
					config.Configurations[i].Types.EnvVar.Value = jsonConfiguration.Types.Json.Value
				}
			}
		}
	}
	if err := validateConfig(config); err != nil {
		return err
	}
	return nil
}

func validateConfig(config *Config) error {
	for _, configuration := range config.Configurations {
		if configuration.Types.Json.Value == "" {
			return errors.New(configuration.Types.EnvVar.Name + "|" + configuration.Types.Json.Name + " is required, either as an environment variable or in the config file.")
		}
		if configuration.Types.EnvVar.Value != configuration.Types.Json.Value {
			return errors.New("Environment Variable " + configuration.Types.EnvVar.Name + ":" + configuration.Types.EnvVar.Value + ", Config file " + configuration.Types.Json.Name + ":" + configuration.Types.Json.Value + " values do not match.")
		}
	}
	return nil
}

func GetConfig() *Config {
	return config
}
