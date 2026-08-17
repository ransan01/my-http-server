package config

import "errors"

type Config struct {
	Address  string
	Port     string
	CertFile string
	KeyFile  string
}

var ConfigEnvVarsToStructMap = map[string]string{
	"MY_API_SERVER_ADDRESS":   "Address",
	"MY_API_SERVER_PORT":      "Port",
	"MY_API_SERVER_CERT_FILE": "CertFile",
	"MY_API_SERVER_KEY_FILE":  "KeyFile",
}

func NewConfig() *Config {
	return &Config{}
}

func LoadConfig() (*Config, error) {
	configEnv := LoadConfigFromEnv()
	configFile, err := LoadConfigFromJsonFile("config.json")
	if err != nil {
		return nil, err
	}
	config := &Config{
		Address:  useConfigValue(configEnv.Address, configFile.Address),
		Port:     useConfigValue(configEnv.Port, configFile.Port),
		CertFile: useConfigValue(configEnv.CertFile, configFile.CertFile),
		KeyFile:  useConfigValue(configEnv.KeyFile, configFile.KeyFile),
	}
	if err := validateConfig(config); err != nil {
		return nil, err
	}
	return config, nil
}

// Environment variable value takes precedence over file value
func useConfigValue(configEnvValue string, configFileValue string) string {
	if configEnvValue != "" {
		return configEnvValue
	}
	return configFileValue
}

func validateConfig(config *Config) error {
	if config.Address == "" {
		return errors.New("address is required")
	}
	if config.Port == "" {
		return errors.New("port is required")
	}
	if config.CertFile == "" {
		return errors.New("cert file is required")
	}
	if config.KeyFile == "" {
		return errors.New("key file is required")
	}
	return nil
}

func LoadConfigOnlyFromEnv() (*Config, error) {
	config := LoadConfigFromEnv()
	if err := validateConfig(config); err != nil {
		return nil, err
	}
	return config, nil
}

func LoadConfigOnlyFromJsonFile() (*Config, error) {
	configFile, err := LoadConfigFromJsonFile("config.json")
	if err != nil {
		return nil, err
	}
	if err := validateConfig(configFile); err != nil {
		return nil, err
	}
	return configFile, nil
}
