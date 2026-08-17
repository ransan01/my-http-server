package config

import (
	"os"
	"reflect"
)

func LoadConfigFromEnv() *Config {
	config := NewConfig()
	for key, value := range ConfigEnvVarsToStructMap {
		envValue := os.Getenv(key)
		if envValue == "" {
			continue
		}
		reflect.ValueOf(config).Elem().FieldByName(value).SetString(envValue)
	}
	return config
}
